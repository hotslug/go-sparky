package installer

import (
	"os"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
	"github.com/hotslug/go-sparky/internal/templates"
)

// InstallHusky installs Husky, initializes hooks, and writes lint-staged config.
func InstallHusky() error {
	logger.Step("Installing Husky and lint-staged")
	if err := runner.Run("pnpm", "install", "-D", "husky@latest", "lint-staged@latest"); err != nil {
		return err
	}

	if err := runner.Run("pnpm", "dlx", "husky-init", "--no-install"); err != nil {
		return err
	}

	if err := os.WriteFile(".lintstagedrc", []byte(templates.LintStagedConfig()), 0o644); err != nil {
		return err
	}

	if err := os.MkdirAll(".husky", 0o755); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(".husky", "pre-commit"), []byte(templates.HuskyPreCommit()), 0o755)
}
