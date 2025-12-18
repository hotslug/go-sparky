package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallMantine installs Mantine dependencies.
func InstallMantine() error {
	spin := logger.StartSpinner("Installing Mantine")
	if err := runner.RunQuiet("pnpm", "install", "@mantine/core@latest", "@mantine/hooks@latest"); err != nil {
		spin("Failed to install Mantine")
		return err
	}
	spin("Installed Mantine")
	return nil
}
