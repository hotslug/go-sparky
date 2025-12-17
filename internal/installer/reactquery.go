package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallReactQuery installs TanStack Query dependencies.
func InstallReactQuery() error {
	spin := logger.StartSpinner("Installing TanStack Query")
	if err := runner.RunQuiet("pnpm", "install", "@tanstack/react-query@latest", "@tanstack/react-query-devtools@latest"); err != nil {
		spin("Failed to install TanStack Query")
		return err
	}
	spin("Installed TanStack Query")
	return nil
}
