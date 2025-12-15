package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallReactQuery installs TanStack Query dependencies.
func InstallReactQuery() error {
	logger.Step("Installing TanStack Query")
	return runner.Run("pnpm", "install", "@tanstack/react-query@latest", "@tanstack/react-query-devtools@latest")
}
