package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/hotslug/go-sparky/internal/installer"
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/runner"
	"github.com/hotslug/go-sparky/internal/version"
	"github.com/spf13/cobra"
)

func newViteSetupCmd() *cobra.Command {
	var (
		flagMantine      bool
		flagNoTailwind   bool
		flagNoReactQuery bool
		flagNoZustand    bool
		flagNoEslint     bool
		flagNoPrettier   bool
		flagNoHusky      bool
		flagStyled       bool
		flagNoFramer     bool
		flagDocker       bool
		flagVercel       bool
		flagNetlify      bool
		flagStorybook    bool
		flagBackend      bool
	)

	cmd := &cobra.Command{
		Use:   "vite-setup [name]",
		Short: "Create a new React app powered by Vite",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]

			logger.PrintBanner()

			if flagStyled && !flagMantine {
				return fmt.Errorf("--styled requires --mantine")
			}

			p := plan.Plan{
				Name:       projectName,
				Bundler:    plan.BundlerVite,
				Mantine:    flagMantine,
				Tailwind:   !flagNoTailwind,
				ReactQuery: !flagNoReactQuery,
				Zustand:    !flagNoZustand,
				Eslint:     !flagNoEslint,
				Prettier:   !flagNoPrettier,
				Husky:      !flagNoHusky,
				StyledApp:  flagStyled && flagMantine,
				Framer:     !flagNoFramer,
				Docker:     flagDocker,
				Vercel:     flagVercel,
				Netlify:    flagNetlify,
				Storybook:  flagStorybook,
				Backend:    flagBackend,
			}

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if flagBackend {
				if _, err := exec.LookPath("bun"); err != nil {
					return fmt.Errorf("bun not found: %w", err)
				}
			}

			if err := version.CheckNodeVersion(); err != nil {
				return err
			}

			if _, err := os.Stat(projectName); err == nil {
				return fmt.Errorf("Project directory \x1b[38;2;255;185;0m%s\x1b[0m already exists. Please choose a different name.", projectName)
			} else if !errors.Is(err, os.ErrNotExist) {
				return err
			}

			logger.Step("Creating project directory")
			if err := os.Mkdir(projectName, 0o755); err != nil {
				return err
			}

			if err := os.Chdir(projectName); err != nil {
				return err
			}

			spin := logger.StartSpinner("Scaffolding with Vite (React + TypeScript)")
			// Set CI to keep create-vite non-interactive.
			// Note: pnpm doesn't require an extra `--` separator to forward args to the starter.
			if err := runner.RunQuietEnv("pnpm", map[string]string{"CI": "1"}, "create", "vite@latest", ".", "--template", "react-ts"); err != nil {
				spin("Failed to scaffold project")
				return err
			}
			spin("Scaffolded Vite project")

			if err := installer.InstallViteReactPlugin(p); err != nil {
				return err
			}

			if p.Mantine {
				if err := installer.InstallMantine(p); err != nil {
					return err
				}
			}

			if p.Framer {
				if err := installer.InstallFramerMotion(p); err != nil {
					return err
				}
			}

			if p.Tailwind {
				if err := installer.InstallTailwind(p); err != nil {
					return err
				}
			}

			if p.ReactQuery {
				if err := installer.InstallReactQuery(p); err != nil {
					return err
				}
			}

			if p.Zustand {
				if err := installer.InstallZustand(p); err != nil {
					return err
				}
			}

			if p.Eslint {
				if err := installer.InstallESLint(p); err != nil {
					return err
				}
			}

			if p.Prettier {
				if err := installer.InstallPrettier(p); err != nil {
					return err
				}
			}

			if p.Husky {
				if err := installer.InstallHusky(p); err != nil {
					return err
				}
			}

			if p.Storybook {
				if err := installer.InstallStorybook(p); err != nil {
					return err
				}
			}

			spin = logger.StartSpinner("Finalizing templates")
			if err := installer.WriteConfigFiles(p); err != nil {
				spin("Failed to finalize templates")
				return err
			}

			if p.Mantine {
				if err := installer.WritePostCSSConfig(); err != nil {
					spin("Failed to finalize templates")
					return err
				}
			}

			if err := installer.WriteAppFiles(p); err != nil {
				spin("Failed to finalize templates")
				return err
			}

			if p.Backend {
				if err := installer.WriteBunBackend(); err != nil {
					spin("Failed to finalize templates")
					return err
				}
			}

			if p.Storybook {
				if err := installer.WriteStorybookConfig(p, true); err != nil {
					spin("Failed to finalize templates")
					return err
				}
			}
			spin("Templates ready")

			if p.Docker {
				if err := installer.WriteDockerArtifacts(p); err != nil {
					return err
				}
			}

			if p.Vercel {
				if err := installer.WriteVercelConfig(p); err != nil {
					return err
				}
			}

			if p.Netlify {
				if err := installer.WriteNetlifyConfig(p); err != nil {
					return err
				}
			}

			spin = logger.StartSpinner("Installing dependencies")
			if err := runner.RunQuiet(p.PackageManager(), "install"); err != nil {
				spin("Failed to install dependencies")
				return err
			}
			spin("Installed dependencies")

			if err := installer.CreateInitialCommitIfMissing("chore: scaffold project"); err != nil {
				return err
			}

			logger.Info("\n⚡ Go Sparky!\n\n→ cd " + projectName + "\n→ " + p.PackageManager() + " dev\n\n⚡ Edit src/App.tsx to begin")

			logger.Info("\nStarting dev server (press Ctrl+C to stop)...")
			if err := runner.Run(p.PackageManager(), "dev"); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&flagMantine, "mantine", false, "Install Mantine")
	cmd.Flags().BoolVar(&flagNoTailwind, "no-tailwind", false, "Skip Tailwind (default installs)")
	cmd.Flags().BoolVar(&flagNoReactQuery, "no-react-query", false, "Skip TanStack Query (default installs)")
	cmd.Flags().BoolVar(&flagNoZustand, "no-zustand", false, "Skip Zustand (default installs)")
	cmd.Flags().BoolVar(&flagNoEslint, "no-eslint", false, "Skip ESLint (default installs)")
	cmd.Flags().BoolVar(&flagNoPrettier, "no-prettier", false, "Skip Prettier (default installs)")
	cmd.Flags().BoolVar(&flagNoHusky, "no-husky", false, "Skip Husky + lint-staged (default installs)")
	cmd.Flags().BoolVar(&flagStyled, "styled", false, "Use styled App template (requires mantine)")
	cmd.Flags().BoolVar(&flagNoFramer, "no-framer-motion", false, "Skip Framer Motion (default installs)")
	cmd.Flags().BoolVar(&flagDocker, "docker", false, "Add Dockerfile and docker-compose.yml")
	cmd.Flags().BoolVar(&flagVercel, "vercel", false, "Add Vercel static build config")
	cmd.Flags().BoolVar(&flagNetlify, "netlify", false, "Add Netlify deploy config")
	cmd.Flags().BoolVar(&flagStorybook, "storybook", false, "Add Storybook config and dependencies")
	cmd.Flags().BoolVar(&flagBackend, "backend", false, "Add a Bun backend in /backend")

	return cmd
}

func newNewCmd() *cobra.Command {
	cmd := newViteSetupCmd()
	cmd.Use = "new [name]"
	cmd.Short = "Deprecated: use vite-setup to create a Vite project"
	cmd.Deprecated = "use vite-setup instead"

	run := cmd.RunE
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		logger.Warning("`new` is deprecated; use `vite-setup` instead.")
		if run != nil {
			return run(cmd, args)
		}
		return nil
	}

	return cmd
}
