package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallViteReactPlugin ensures @vitejs/plugin-react is installed to match our Vite config template.
func InstallViteReactPlugin() error {
	spin := logger.StartSpinner("Installing Vite React plugin")
	if err := runner.RunQuiet("pnpm", "install", "-D", "@vitejs/plugin-react@latest"); err != nil {
		spin("Failed to install Vite React plugin")
		return err
	}
	spin("Installed Vite React plugin")
	return nil
}
