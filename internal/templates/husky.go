package templates

import "github.com/hotslug/go-sparky/internal/plan"

// LintStagedConfig returns the .lintstagedrc template.
func LintStagedConfig(p plan.Plan) string {
	commandPrefix := "pnpm"
	if p.IsBun() {
		commandPrefix = "bun run"
	}

	return `{
  "*.{js,jsx,ts,tsx}": ["` + commandPrefix + ` eslint --fix"],
  "*.{js,jsx,ts,tsx,css,md,json}": ["` + commandPrefix + ` prettier --write"]
}
`
}

// HuskyPreCommit returns the pre-commit hook content.
func HuskyPreCommit(p plan.Plan) string {
	commandPrefix := "pnpm"
	if p.IsBun() {
		commandPrefix = "bun run"
	}

	return `#!/bin/sh
. "$(dirname "$0")/_/husky.sh"

` + commandPrefix + ` lint-staged
`
}
