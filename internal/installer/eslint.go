package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
	"github.com/hotslug/go-sparky/internal/templates"
)

// InstallESLint installs ESLint dependencies and config.
func InstallESLint() error {
	spin := logger.StartSpinner("Installing ESLint")
	if err := runner.RunQuiet("pnpm", "install", "-D",
		"eslint@latest",
		"eslint-plugin-react-x@latest",
		"eslint-plugin-react-dom@latest",
		"@typescript-eslint/parser@latest",
	); err != nil {
		spin("Failed to install ESLint")
		return err
	}
	spin("Installed ESLint")

	return os.WriteFile("eslint.config.js", []byte(templates.EslintConfig()), 0o644)
}
