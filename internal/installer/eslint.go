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
		"@eslint/js@latest",
		"@typescript-eslint/parser@latest",
		"@typescript-eslint/eslint-plugin@latest",
		"@tanstack/eslint-plugin-query@latest",
		"eslint-import-resolver-typescript@latest",
		"eslint-plugin-react@latest",
		"eslint-plugin-react-hooks@latest",
		"eslint-plugin-jsx-a11y@latest",
		"eslint-plugin-import@latest",
		"eslint-plugin-unicorn@latest",
		"eslint-plugin-prettier@latest",
		"eslint-config-prettier@latest",
	); err != nil {
		spin("Failed to install ESLint")
		return err
	}
	spin("Installed ESLint")

	return os.WriteFile("eslint.config.js", []byte(templates.EslintConfig()), 0o644)
}
