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
	"github.com/spf13/cobra"
)

func newNewCmd() *cobra.Command {
	var (
		flagMantine      bool
		flagNoTailwind   bool
		flagNoReactQuery bool
		flagNoEslint     bool
		flagNoPrettier   bool
		flagNoHusky      bool
		flagStyled       bool
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
				Eslint:     !flagNoEslint,
				Prettier:   !flagNoPrettier,
				Husky:      !flagNoHusky,
				StyledApp:  flagStyled && flagMantine,
				Docker:     flagDocker,
				Vercel:     flagVercel,
				Netlify:    flagNetlify,
			}

			if _, err := exec.LookPath("pnpm"); err != nil {
				return fmt.Errorf("pnpm not found: %w", err)
			}

			if _, err := os.Stat(projectName); err == nil {
				return fmt.Errorf("project directory %s already exists", projectName)
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

			logger.Step("Scaffolding with Vite (React + TypeScript)")
			logger.Info("Command: pnpm create vite@latest . -- --template react-ts --no-interactive --install=false --rolldown=false --packageManager pnpm --skipGit")
			// Forward flags after -- and set CI to keep create-vite non-interactive.
			if err := runner.RunEnv("pnpm", map[string]string{"CI": "1"}, "create", "vite@latest", ".", "--", "--template", "react-ts", "--no-interactive", "--install=false", "--rolldown=false", "--packageManager", "pnpm", "--skipGit"); err != nil {
				return err
			}

			if err := installer.InstallViteReactPlugin(); err != nil {
				return err
			}

			if p.Mantine {
				if err := installer.InstallMantine(); err != nil {
					return err
				}
			}

			if err := installer.InstallFramerMotion(); err != nil {
				return err
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

			spin := logger.StartSpinner("Finalizing templates")
			if err := installer.WriteViteConfig(p.Tailwind); err != nil {
				spin("Failed to finalize templates")
				return err
			}

			if p.Mantine {
				if err := installer.WritePostCSSConfig(p.Tailwind); err != nil {
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

			logger.Step("Installing dependencies")
			if err := runner.Run("pnpm", "install"); err != nil {
				return err
			}

			logger.Info("Starting dev server with pnpm dev (press Ctrl+C to stop)")
			if err := runner.Run("pnpm", "dev"); err != nil {
				return err
			}

			logger.Info("✨ Sparky is ready!\n\n→ cd " + projectName + "\n→ pnpm dev\n\nEdit src/App.tsx to begin ⚡")
			return nil
		},
	}

	cmd.Flags().BoolVar(&flagMantine, "mantine", false, "Install Mantine")
	cmd.Flags().BoolVar(&flagNoTailwind, "no-tailwind", false, "Skip Tailwind (default installs)")
	cmd.Flags().BoolVar(&flagNoReactQuery, "no-react-query", false, "Skip TanStack Query (default installs)")
	cmd.Flags().BoolVar(&flagNoEslint, "no-eslint", false, "Skip ESLint (default installs)")
	cmd.Flags().BoolVar(&flagNoPrettier, "no-prettier", false, "Skip Prettier (default installs)")
	cmd.Flags().BoolVar(&flagNoHusky, "no-husky", false, "Skip Husky + lint-staged (default installs)")
	cmd.Flags().BoolVar(&flagStyled, "styled", false, "Use styled App template (requires mantine)")
	cmd.Flags().BoolVar(&flagDocker, "docker", false, "Add Dockerfile and docker-compose.yml")
	cmd.Flags().BoolVar(&flagVercel, "vercel", false, "Add Vercel static build config")
	cmd.Flags().BoolVar(&flagNetlify, "netlify", false, "Add Netlify deploy config")

	return cmd
}
