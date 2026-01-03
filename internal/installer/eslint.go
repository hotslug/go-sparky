package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/templates"
)

// InstallESLint installs ESLint dependencies and config.
func InstallESLint(p plan.Plan) error {
	spin := logger.StartSpinner("Installing ESLint")
	if err := addDependencies(p, true,
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

	return os.WriteFile("eslint.config.js", []byte(templates.EslintConfig(p)), 0o644)
}

// WriteESLintStrict rewrites eslint.config.js with the default strict config.
func WriteESLintStrict(p plan.Plan) error {
	return os.WriteFile("eslint.config.js", []byte(templates.EslintConfig(p)), 0o644)
}

// WriteESLintRelaxed rewrites eslint.config.js with a looser preset.
func WriteESLintRelaxed(p plan.Plan) error {
	return os.WriteFile("eslint.config.js", []byte(templates.EslintConfigRelaxed(p)), 0o644)
}
