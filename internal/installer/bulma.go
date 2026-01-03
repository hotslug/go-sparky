package installer

import (
	"bytes"
	"os"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
)

// InstallBulma installs Bulma CSS.
func InstallBulma(p plan.Plan) error {
	spin := logger.StartSpinner("Installing Bulma")
	if err := addDependencies(p, false, "bulma@latest"); err != nil {
		spin("Failed to install Bulma")
		return err
	}
	spin("Installed Bulma")
	return nil
}

// RemoveBulma uninstalls Bulma CSS.
func RemoveBulma(p plan.Plan) error {
	spin := logger.StartSpinner("Removing Bulma")
	if err := removeDependencies(p, false, "bulma"); err != nil {
		spin("Failed to remove Bulma")
		return err
	}
	spin("Removed Bulma")
	return nil
}

// EnsureBulmaImport prepends an @import to the given CSS file if not already present.
func EnsureBulmaImport(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if bytes.Contains(data, []byte("bulma/css/bulma.min.css")) {
		return nil
	}

	content := []byte("@import 'bulma/css/bulma.min.css';\n" + string(data))
	return os.WriteFile(path, content, 0o644)
}
