package cmd

import (
	"fmt"
	"os/exec"

	"github.com/hotslug/go-sparky/internal/installer"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/version"
)

func detectBundlerPlan() (plan.Plan, error) {
	bundler, err := installer.DetectBundler()
	if err != nil {
		return plan.Plan{}, err
	}

	pkgMgr := "pnpm"
	if bundler == plan.BundlerBun {
		pkgMgr = "bun"
	}

	if _, err := exec.LookPath(pkgMgr); err != nil {
		return plan.Plan{}, fmt.Errorf("%s not found: %w", pkgMgr, err)
	}

	if bundler == plan.BundlerVite {
		if err := version.CheckNodeVersion(); err != nil {
			return plan.Plan{}, err
		}
	}

	return plan.Plan{Bundler: bundler}, nil
}

func mainEntryFilename(p plan.Plan) string {
	if p.IsBun() {
		return "frontend.tsx"
	}
	return "main.tsx"
}

func storybookCommand(p plan.Plan) string {
	if p.IsBun() {
		return "bun run storybook dev -p 6006"
	}
	return "pnpm storybook dev -p 6006"
}
