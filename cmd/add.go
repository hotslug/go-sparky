package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/installer"
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/runner"
	"github.com/hotslug/go-sparky/internal/version"
	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add optional stacks to an existing project",
	}

	cmd.AddCommand(newAddMantineCmd())
	cmd.AddCommand(newAddReactQueryCmd())
	cmd.AddCommand(newAddZustandCmd())
	cmd.AddCommand(newAddDockerCmd())
	cmd.AddCommand(newAddVercelCmd())
	cmd.AddCommand(newAddNetlifyCmd())
	cmd.AddCommand(newAddFramerMotionCmd())
	cmd.AddCommand(newAddShadcnCmd())
	cmd.AddCommand(newAddBulmaCmd())
	cmd.AddCommand(newAddStorybookCmd())
	return cmd
}

func newAddMantineCmd() *cobra.Command {
	var flagStyled bool

	cmd := &cobra.Command{
		Use:   "mantine",
		Short: "Install Mantine and wire it into main.tsx without touching App.tsx",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if flagStyled {
				return fmt.Errorf("--styled is only available during scaffolding (go-sparky new --mantine --styled). The add command intentionally leaves src/App.tsx alone so you can keep your existing UI.")
			}

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			mainPath := filepath.Join("src", "main.tsx")
			mainContent, err := os.ReadFile(mainPath)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("src/main.tsx not found. Run this in a Vite React project")
				}
				return err
			}

			if err := installer.InstallMantine(); err != nil {
				return err
			}

			if err := installer.WritePostCSSConfig(); err != nil {
				return err
			}

			if bytes.Contains(mainContent, []byte("MantineProvider")) {
				logger.Info("\nMantineProvider already detected in src/main.tsx; leaving the file unchanged.")
				logger.Info("\nMantine packages installed and PostCSS configured. App.tsx was not modified.")
				return nil
			}

			p := plan.Plan{
				Mantine:    true,
				ReactQuery: installer.HasReactQueryDependency(),
			}

			if err := installer.WriteMainFile(p); err != nil {
				return err
			}

			logger.Info("\nMantine added. src/main.tsx updated with MantineProvider. App.tsx left untouched.")
			return nil
		},
	}

	cmd.Flags().BoolVar(&flagStyled, "styled", false, "Not supported: styled template only applies during scaffolding")
	return cmd
}

func newAddReactQueryCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "react-query",
		Short: "Install TanStack Query and wrap providers in main.tsx (App.tsx untouched)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			mainPath := filepath.Join("src", "main.tsx")
			mainContent, err := os.ReadFile(mainPath)
			if err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("src/main.tsx not found. Run this in a Vite React project")
				}
				return err
			}

			if err := installer.InstallReactQuery(); err != nil {
				return err
			}

			if bytes.Contains(mainContent, []byte("QueryClientProvider")) {
				logger.Info("\nQueryClientProvider already detected in src/main.tsx; leaving the file unchanged.")
				logger.Info("\nReact Query packages installed. App.tsx was not modified.")
				return nil
			}

			p := plan.Plan{
				ReactQuery: true,
				Mantine:    installer.HasMantineDependency(),
			}

			if err := installer.WriteMainFile(p); err != nil {
				return err
			}

			logger.Info("\nReact Query added. src/main.tsx updated with QueryClientProvider. App.tsx left untouched.")
			return nil
		},
	}
}

func newAddZustandCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "zustand",
		Short: "Install Zustand and add a starter store (App.tsx untouched)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			if err := installer.InstallZustand(); err != nil {
				return err
			}

			created, err := installer.WriteZustandStoreIfMissing()
			if err != nil {
				return err
			}

			if created {
				logger.Info("\nZustand installed. Added src/stores/useSparkyStore.ts. App.tsx left untouched.")
				return nil
			}

			logger.Info("\nZustand installed. src/stores/useSparkyStore.ts already exists; left untouched. App.tsx left untouched.")
			return nil
		},
	}
}

func newAddDockerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "docker",
		Short: "Add Dockerfile and docker-compose.yml",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if err := installer.WriteDockerArtifacts(); err != nil {
				return err
			}

			logger.Info("\nDocker artifacts written (Dockerfile, docker-compose.yml).")
			return nil
		},
	}
}

func newAddVercelCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "vercel",
		Short: "Add Vercel static build config",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if err := installer.WriteVercelConfig(); err != nil {
				return err
			}

			logger.Info("\nvercel.json written.")
			return nil
		},
	}
}

func newAddNetlifyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "netlify",
		Short: "Add Netlify deploy config",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if err := installer.WriteNetlifyConfig(); err != nil {
				return err
			}

			logger.Info("\nnetlify.toml written.")
			return nil
		},
	}
}

func newAddFramerMotionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "framer-motion",
		Short: "Install Framer Motion",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			if err := installer.InstallFramerMotion(); err != nil {
				return err
			}

			logger.Info("\nFramer Motion installed. App.tsx left untouched.")
			return nil
		},
	}
}

func newAddStorybookCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "storybook",
		Short: "Install Storybook with Vite + React defaults",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			if installer.HasStorybookConfig() {
				logger.Warning("\n.storybook already exists; leaving your Storybook config unchanged.")
				logger.Info("\nStart it with `pnpm storybook dev -p 6006` or update your existing config manually.")
				return nil
			}

			indexCSSExists := true
			if _, err := os.Stat(filepath.Join("src", "index.css")); err != nil {
				if os.IsNotExist(err) {
					indexCSSExists = false
				} else {
					return err
				}
			}

			if err := installer.InstallStorybook(); err != nil {
				return err
			}

			if err := installer.WriteStorybookConfig(indexCSSExists); err != nil {
				return err
			}

			if !indexCSSExists {
				logger.Warning("\nsrc/index.css not found; update .storybook/preview.ts to import your global styles if needed.")
			}

			logger.Info("\nStorybook added. Start it with `pnpm storybook dev -p 6006`.")
			return nil
		},
	}
}

func newAddShadcnCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "shadcn",
		Short: "Run shadcn-ui init (interactive) on top of Tailwind",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			indexCSS := filepath.Join("src", "index.css")
			indexExists := true
			if _, err := os.Stat(indexCSS); err != nil {
				if os.IsNotExist(err) {
					indexExists = false
				} else {
					return err
				}
			}

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			if !installer.HasTailwind() {
				return fmt.Errorf("Tailwind not detected. shadcn/ui requires Tailwind; rerun with Tailwind enabled or install Tailwind first")
			}

			if _, err := os.Stat("components.json"); err == nil {
				logger.Warning("\ncomponents.json already exists; shadcn/ui looks initialized. Skipping init.")
				logger.Info("\nUse `pnpm dlx shadcn-ui@latest add <component>` to add components.")
				return nil
			} else if err != nil && !os.IsNotExist(err) {
				return err
			}

			logger.Info("\nRunning shadcn-ui init (you'll see prompts for theme/config)...")
			if err := runner.Run("pnpm", "dlx", "shadcn-ui@latest", "init"); err != nil {
				return err
			}

			if !indexExists {
				logger.Warning("\nshadcn-ui initialized, but src/index.css was not found; ensure your Tailwind entry CSS exists and is wired in your Vite project.")
			}

			logger.Info("\nshadcn-ui initialized. Add components with `pnpm dlx shadcn-ui@latest add button card input ...`")
			return nil
		},
	}
}

func newAddBulmaCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bulma",
		Short: "Install Bulma CSS (App.tsx untouched; import it in your CSS)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			indexCSS := filepath.Join("src", "index.css")
			indexExists := true
			if _, err := os.Stat(indexCSS); err != nil {
				if os.IsNotExist(err) {
					indexExists = false
				} else {
					return err
				}
			}

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat("package.json"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("package.json not found. Run this inside your existing app directory")
				}
				return err
			}

			if err := installer.InstallBulma(); err != nil {
				return err
			}

			if !indexExists {
				logger.Warning("\nBulma installed, but src/index.css was not found; add `@import 'bulma/css/bulma.min.css';` to your global CSS manually.")
				return nil
			}

			if err := installer.EnsureBulmaImport(indexCSS); err != nil {
				logger.Warning("\nBulma installed but could not update " + indexCSS + ": " + err.Error())
				logger.Info("Add `@import 'bulma/css/bulma.min.css';` to your global CSS (after Tailwind directives if you want Tailwind to win). App.tsx left untouched.")
				return nil
			}

			logger.Info("\nBulma installed and @import added to src/index.css (placed at the top). App.tsx left untouched.")
			return nil
		},
	}
}
