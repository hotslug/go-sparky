package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallViteReactPlugin ensures @vitejs/plugin-react is installed to match our Vite config template.
func InstallViteReactPlugin() error {
	logger.Step("Ensuring @vitejs/plugin-react is installed")
	return runner.Run("pnpm", "install", "-D", "@vitejs/plugin-react@latest")
}
