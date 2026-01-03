package installer

import (
	"os"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/runner"
	"github.com/hotslug/go-sparky/internal/templates"
)

// InstallHusky installs Husky, initializes hooks, and writes lint-staged config.
func InstallHusky(p plan.Plan) error {
	if _, err := os.Stat(".git"); err != nil {
		if os.IsNotExist(err) {
			spin := logger.StartSpinner("Initializing git repository")
			if err := runner.RunQuiet("git", "init", "-b", "main"); err != nil {
				spin("Failed to initialize git repository")
				return err
			}
			spin("Initialized git repository")
		} else {
			return err
		}
	}

	spin := logger.StartSpinner("Installing Husky and lint-staged")
	packages := []string{"husky@latest", "lint-staged@latest"}
	if p.IsBun() {
		packages = append(packages, "husky-init@latest")
	}

	if err := addDependencies(p, true, packages...); err != nil {
		spin("Failed to install Husky and lint-staged")
		return err
	}

	if p.IsBun() {
		if err := runner.RunQuiet("bun", "run", "husky-init", "--", "--no-install"); err != nil {
			spin("Failed to initialize Husky")
			return err
		}
	} else {
		if err := runner.RunQuiet("pnpm", "dlx", "husky-init", "--no-install"); err != nil {
			spin("Failed to initialize Husky")
			return err
		}
	}
	spin("Installed Husky and lint-staged")

	if err := os.WriteFile(".lintstagedrc", []byte(templates.LintStagedConfig(p)), 0o644); err != nil {
		return err
	}

	if err := os.MkdirAll(".husky", 0o755); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(".husky", "pre-commit"), []byte(templates.HuskyPreCommit(p)), 0o755)
}
