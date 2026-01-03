package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/installer"
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/templates"
	"github.com/spf13/cobra"
)

func newRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove optional stacks from an existing project",
	}

	cmd.AddCommand(newRemoveMantineCmd())
	cmd.AddCommand(newRemoveReactQueryCmd())
	cmd.AddCommand(newRemoveZustandCmd())
	cmd.AddCommand(newRemoveDockerCmd())
	cmd.AddCommand(newRemoveVercelCmd())
	cmd.AddCommand(newRemoveNetlifyCmd())
	cmd.AddCommand(newRemoveFramerMotionCmd())
	cmd.AddCommand(newRemoveBulmaCmd())
	return cmd
}

func newRemoveMantineCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mantine",
		Short: "Uninstall Mantine, delete PostCSS config, and unwrap MantineProvider in main.tsx",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			p, err := detectBundlerPlan()
			if err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			mainPath := filepath.Join("src", mainEntryFilename(p))
			mainContent, err := os.ReadFile(mainPath)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("%s not found. Run this in a go-sparky project", mainPath)
				}
				return err
			}

			if err := installer.RemoveMantine(p); err != nil {
				return err
			}

			if err := installer.DeletePostCSSConfigIfOwned(); err != nil {
				return err
			}

			if !bytes.Contains(mainContent, []byte("MantineProvider")) {
				logger.Info("\nMantineProvider not found in src/main.tsx; leaving the file unchanged.")
				logger.Info("\nMantine packages removed. App.tsx was not modified.")
				return nil
			}

			p.ReactQuery = installer.HasReactQueryDependency()

			if err := installer.WriteMainFile(p, mainEntryFilename(p)); err != nil {
				return err
			}

			logger.Info("\nMantine removed. src/main.tsx updated to remove MantineProvider. App.tsx left untouched.")
			return nil
		},
	}
}

func newRemoveReactQueryCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "react-query",
		Short: "Uninstall TanStack Query and unwrap providers in main.tsx",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			p, err := detectBundlerPlan()
			if err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			mainPath := filepath.Join("src", mainEntryFilename(p))
			mainContent, err := os.ReadFile(mainPath)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("%s not found. Run this in a go-sparky project", mainPath)
				}
				return err
			}

			if err := installer.RemoveReactQuery(p); err != nil {
				return err
			}

			if !bytes.Contains(mainContent, []byte("QueryClientProvider")) {
				logger.Info("\nQueryClientProvider not found in src/main.tsx; leaving the file unchanged.")
				logger.Info("\nReact Query packages removed. App.tsx was not modified.")
				return nil
			}

			p.Mantine = installer.HasMantineDependency()

			if err := installer.WriteMainFile(p, mainEntryFilename(p)); err != nil {
				return err
			}

			logger.Info("\nReact Query removed. src/main.tsx updated to remove QueryClientProvider. App.tsx left untouched.")
			return nil
		},
	}
}

func newRemoveZustandCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "zustand",
		Short: "Uninstall Zustand and remove the starter store if unmodified",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			p, err := detectBundlerPlan()
			if err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			appPath := filepath.Join("src", "App.tsx")
			appContent, err := os.ReadFile(appPath)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("src/App.tsx not found. Run this in a go-sparky project")
				}
				return err
			}

			appMatchesGenerated := string(appContent) == templates.AppTemplate(plan.Plan{Zustand: true})
			appUsesStore := bytes.Contains(appContent, []byte("useSparkyStore"))

			if err := installer.RemoveZustand(p); err != nil {
				return err
			}

			switch {
			case appMatchesGenerated:
				if err := installer.WriteAppFile(plan.Plan{}); err != nil {
					return err
				}
				if err := installer.DeleteZustandStoreIfOwned(); err != nil {
					return err
				}
				logger.Info("\nZustand removed. src/App.tsx reset to the basic template and demo store deleted.")
			case !appUsesStore:
				if err := installer.DeleteZustandStoreIfOwned(); err != nil {
					return err
				}
				logger.Info("\nZustand removed. App.tsx left untouched.")
			default:
				logger.Warning("\nZustand removed, but src/App.tsx still references useSparkyStore; update your state to avoid missing imports.")
			}

			return nil
		},
	}
}

func newRemoveDockerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "docker",
		Short: "Delete generated Docker artifacts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if err := installer.DeleteDockerArtifacts(); err != nil {
				return err
			}

			logger.Info("\nDocker artifacts removed if they matched the generated content.")
			return nil
		},
	}
}

func newRemoveVercelCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "vercel",
		Short: "Delete generated vercel.json if unmodified",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if err := installer.DeleteVercelConfig(); err != nil {
				return err
			}

			logger.Info("\nvercel.json removed if it matched the generated content.")
			return nil
		},
	}
}

func newRemoveNetlifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "netlify",
		Short: "Delete generated netlify.toml if unmodified",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if err := installer.DeleteNetlifyConfig(); err != nil {
				return err
			}

			logger.Info("\nnetlify.toml removed if it matched the generated content.")
			return nil
		},
	}
}

func newRemoveFramerMotionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "framer-motion",
		Short: "Uninstall Framer Motion",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			p, err := detectBundlerPlan()
			if err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			if err := installer.RemoveFramerMotion(p); err != nil {
				return err
			}

			logger.Info("\nFramer Motion removed. App.tsx left untouched.")
			return nil
		},
	}
}

func newRemoveBulmaCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bulma",
		Short: "Uninstall Bulma CSS",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			p, err := detectBundlerPlan()
			if err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			if err := installer.RemoveBulma(p); err != nil {
				return err
			}

			logger.Info("\nBulma removed. App.tsx and CSS files were not modified; remove any Bulma @import you added.")
			return nil
		},
	}
}
