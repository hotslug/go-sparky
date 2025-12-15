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
		flagMantine    bool
		flagTailwind   bool
		flagReactQuery bool
		flagEslint     bool
		flagPrettier   bool
		flagHusky      bool
		flagStyled     bool
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
				Tailwind:   flagTailwind,
				ReactQuery: flagReactQuery,
				Eslint:     flagEslint,
				Prettier:   flagPrettier,
				Husky:      flagHusky,
				StyledApp:  flagStyled && flagMantine,
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
			if err := runner.Run("pnpm", "create", "vite@latest", ".", "--template", "react-ts", "--install"); err != nil {
				return err
			}

			if p.Mantine {
				if err := installer.InstallMantine(); err != nil {
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

			if err := installer.WriteViteConfig(p.Tailwind); err != nil {
				return err
			}

			if err := installer.WritePostCSSConfig(p.Tailwind); err != nil {
				return err
			}

			if err := installer.WriteAppFiles(p); err != nil {
				return err
			}

			logger.Info("✨ Sparky is ready!\n\n→ cd " + projectName + "\n→ pnpm dev\n\nEdit src/App.tsx to begin ⚡")
			return nil
		},
	}

	cmd.Flags().BoolVar(&flagMantine, "mantine", false, "Install Mantine")
	cmd.Flags().BoolVar(&flagTailwind, "tailwind", false, "Install Tailwind")
	cmd.Flags().BoolVar(&flagReactQuery, "react-query", false, "Install TanStack Query")
	cmd.Flags().BoolVar(&flagEslint, "eslint", false, "Install ESLint (react-x + react-dom)")
	cmd.Flags().BoolVar(&flagPrettier, "prettier", false, "Install Prettier")
	cmd.Flags().BoolVar(&flagHusky, "husky", false, "Install Husky + lint-staged")
	cmd.Flags().BoolVar(&flagStyled, "styled", false, "Use styled App template (requires mantine)")

	return cmd
}
