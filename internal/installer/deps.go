package installer

import (
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/runner"
)

func addDependencies(p plan.Plan, dev bool, packages ...string) error {
	if len(packages) == 0 {
		return nil
	}

	args := []string{}
	if p.IsBun() {
		args = append(args, "add")
		if dev {
			args = append(args, "-d")
		}
	} else {
		args = append(args, "install")
		if dev {
			args = append(args, "-D")
		}
	}

	args = append(args, packages...)
	return runner.RunQuiet(p.PackageManager(), args...)
}

func removeDependencies(p plan.Plan, dev bool, packages ...string) error {
	if len(packages) == 0 {
		return nil
	}

	args := []string{"remove"}
	if !p.IsBun() && dev {
		args = append(args, "-D")
	}

	args = append(args, packages...)
	return runner.RunQuiet(p.PackageManager(), args...)
}
