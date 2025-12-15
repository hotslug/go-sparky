package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallMantine installs Mantine dependencies.
func InstallMantine() error {
	logger.Step("Installing Mantine")
	return runner.Run("pnpm", "install", "@mantine/core@latest", "@mantine/hooks@latest")
}
