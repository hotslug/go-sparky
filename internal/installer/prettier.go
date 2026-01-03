package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/templates"
)

// InstallPrettier installs Prettier and writes the config file.
func InstallPrettier(p plan.Plan) error {
	spin := logger.StartSpinner("Installing Prettier")
	if err := addDependencies(p, true,
		"prettier@latest",
		"prettier-plugin-tailwindcss@latest",
		"@ianvs/prettier-plugin-sort-imports@latest",
	); err != nil {
		spin("Failed to install Prettier")
		return err
	}
	spin("Installed Prettier")

	if err := os.WriteFile(".prettierrc", []byte(templates.PrettierConfig()), 0o644); err != nil {
		return err
	}

	return os.WriteFile(".prettierignore", []byte(templates.PrettierIgnore()), 0o644)
}
