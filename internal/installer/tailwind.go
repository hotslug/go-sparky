package installer

import (
	"fmt"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
)

// InstallTailwind installs Tailwind dependencies.
func InstallTailwind(p plan.Plan) error {
	spin := logger.StartSpinner("Installing Tailwind CSS")

	var args []string
	if p.IsVite() {
		args = []string{"tailwindcss@latest", "@tailwindcss/vite@latest"}
	} else if p.IsBun() {
		args = []string{"tailwindcss@^4", "bun-plugin-tailwind@latest"}
	}

	if len(args) == 0 {
		return fmt.Errorf("unknown bundler for Tailwind install")
	}

	if err := addDependencies(p, true, args...); err != nil {
		spin("Failed to install Tailwind CSS")
		return err
	}
	spin("Installed Tailwind CSS")

	if p.IsBun() {
		return WriteBunConfig(p)
	}

	return nil
}
