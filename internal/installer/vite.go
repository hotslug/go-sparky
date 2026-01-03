package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
)

// InstallViteReactPlugin ensures @vitejs/plugin-react is installed to match our Vite config template.
func InstallViteReactPlugin(p plan.Plan) error {
	spin := logger.StartSpinner("Installing Vite React plugin")
	if err := addDependencies(p, true, "@vitejs/plugin-react@latest"); err != nil {
		spin("Failed to install Vite React plugin")
		return err
	}
	spin("Installed Vite React plugin")
	return nil
}
