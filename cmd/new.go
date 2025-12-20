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

func newNewCmd() *cobra.Command {
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
	)

	cmd := &cobra.Command{
		Use:   "new [name]",
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
			}

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
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

			if err := installer.InstallViteReactPlugin(); err != nil {
				return err
			}

			if p.Mantine {
				if err := installer.InstallMantine(); err != nil {
					return err
				}
			}

			if p.Framer {
				if err := installer.InstallFramerMotion(); err != nil {
					return err
				}
			}

			if p.Tailwind {
				if err := installer.InstallTailwind(); err != nil {
					return err
				}
			}

			if p.ReactQuery {
				if err := installer.InstallReactQuery(); err != nil {
					return err
				}
			}

			if p.Zustand {
				if err := installer.InstallZustand(); err != nil {
					return err
				}
			}

			if p.Eslint {
				if err := installer.InstallESLint(); err != nil {
					return err
				}
			}

			if p.Prettier {
				if err := installer.InstallPrettier(); err != nil {
					return err
				}
			}

			if p.Husky {
				if err := installer.InstallHusky(); err != nil {
					return err
				}
			}

			spin = logger.StartSpinner("Finalizing templates")
			if err := installer.WriteViteConfig(p.Tailwind); err != nil {
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
			spin("Templates ready")

			if p.Docker {
				if err := installer.WriteDockerArtifacts(); err != nil {
					return err
				}
			}

			if p.Vercel {
				if err := installer.WriteVercelConfig(); err != nil {
					return err
				}
			}

			if p.Netlify {
				if err := installer.WriteNetlifyConfig(); err != nil {
					return err
				}
			}

			spin = logger.StartSpinner("Installing dependencies")
			if err := runner.RunQuiet("pnpm", "install"); err != nil {
				spin("Failed to install dependencies")
				return err
			}
			spin("Installed dependencies")

			if err := installer.CreateInitialCommitIfMissing("chore: scaffold project"); err != nil {
				return err
			}

			logger.Info("\n⚡ Go Sparky!\n\n→ cd " + projectName + "\n→ pnpm dev\n\n⚡ Edit src/App.tsx to begin")

			logger.Info("\nStarting dev server (press Ctrl+C to stop)...")
			if err := runner.Run("pnpm", "dev"); err != nil {
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

	return cmd
}
