package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallFramerMotion installs Framer Motion dependency.
func InstallFramerMotion() error {
	logger.Step("Installing Framer Motion")
	return runner.Run("pnpm", "install", "framer-motion@latest")
}
