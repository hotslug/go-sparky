package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
	"github.com/hotslug/go-sparky/internal/templates"
)

// InstallPrettier installs Prettier and writes the config file.
func InstallPrettier() error {
	logger.Step("Installing Prettier")
	if err := runner.Run("pnpm", "install", "-D",
		"prettier@latest",
		"prettier-plugin-tailwindcss@latest",
		"@ianvs/prettier-plugin-sort-imports@latest",
	); err != nil {
		return err
	}

	return os.WriteFile(".prettierrc", []byte(templates.PrettierConfig()), 0o644)
}
