package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallTailwind installs Tailwind dependencies.
func InstallTailwind() error {
	spin := logger.StartSpinner("Installing Tailwind CSS")
	if err := runner.RunQuiet("pnpm", "install", "-D", "tailwindcss@latest", "@tailwindcss/vite@latest"); err != nil {
		spin("Failed to install Tailwind CSS")
		return err
	}
	spin("Installed Tailwind CSS")

	return nil
}
