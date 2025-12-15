package templates

// LintStagedConfig returns the .lintstagedrc template.
func LintStagedConfig() string {
	return `{
  "*.{js,jsx,ts,tsx}": ["pnpm eslint --fix"],
  "*.{js,jsx,ts,tsx,css,md,json}": ["pnpm prettier --write"]
}
`
}

// HuskyPreCommit returns the pre-commit hook content.
func HuskyPreCommit() string {
	return `#!/bin/sh
. "$(dirname "$0")/_/husky.sh"

pnpm lint-staged
`
}
