package cmd

import (
	"fmt"
	"os"

	"github.com/hotslug/go-sparky/internal/installer"
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/spf13/cobra"
)

func newLintCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Lint tooling helpers",
	}

	cmd.AddCommand(newLintRelaxCmd())
	cmd.AddCommand(newLintResetCmd())
	return cmd
}

func newLintRelaxCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "relax",
		Short: "Rewrite eslint.config.js with a relaxed preset (no unicorn, import rules off)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			if _, err := os.Stat("eslint.config.js"); err != nil {
				if os.IsNotExist(err) {
					return fmt.Errorf("eslint.config.js not found. Run this inside a scaffolded project with ESLint configured")
				}
				return err
			}

			bundler, err := installer.DetectBundler()
			if err != nil {
				return err
			}

			if err := installer.WriteESLintRelaxed(plan.Plan{Bundler: bundler}); err != nil {
				return err
			}

			logger.Info("\nESLint relaxed: unicorn removed; import ordering/newline rules disabled; core recommended rules kept.")
			return nil
		},
	}
}

func newLintResetCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "reset",
		Short: "Restore eslint.config.js to the default strict preset",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.PrintBanner()

			bundler, err := installer.DetectBundler()
			if err != nil {
				return err
			}

			if err := installer.WriteESLintStrict(plan.Plan{Bundler: bundler}); err != nil {
				return err
			}

			logger.Info("\nESLint reset to default strict config.")
			return nil
		},
	}
}
