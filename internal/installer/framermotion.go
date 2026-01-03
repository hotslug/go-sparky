package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
)

// InstallFramerMotion installs Framer Motion dependency.
func InstallFramerMotion(p plan.Plan) error {
	spin := logger.StartSpinner("Installing Framer Motion")
	if err := addDependencies(p, false, "framer-motion@latest"); err != nil {
		spin("Failed to install Framer Motion")
		return err
	}
	spin("Installed Framer Motion")
	return nil
}

// RemoveFramerMotion uninstalls Framer Motion.
func RemoveFramerMotion(p plan.Plan) error {
	spin := logger.StartSpinner("Removing Framer Motion")
	if err := removeDependencies(p, false, "framer-motion"); err != nil {
		spin("Failed to remove Framer Motion")
		return err
	}
	spin("Removed Framer Motion")
	return nil
}
