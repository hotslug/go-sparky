package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallFramerMotion installs Framer Motion dependency.
func InstallFramerMotion() error {
	spin := logger.StartSpinner("Installing Framer Motion")
	if err := runner.RunQuiet("pnpm", "install", "framer-motion@latest"); err != nil {
		spin("Failed to install Framer Motion")
		return err
	}
	spin("Installed Framer Motion")
	return nil
}
