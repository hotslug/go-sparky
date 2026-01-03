package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
)

// InstallReactQuery installs TanStack Query dependencies.
func InstallReactQuery(p plan.Plan) error {
	spin := logger.StartSpinner("Installing TanStack Query")
	if err := addDependencies(p, false, "@tanstack/react-query@latest", "@tanstack/react-query-devtools@latest"); err != nil {
		spin("Failed to install TanStack Query")
		return err
	}
	spin("Installed TanStack Query")
	return nil
}
