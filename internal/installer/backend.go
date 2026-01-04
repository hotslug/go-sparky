package installer

import (
	"os"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/templates"
)

// WriteBunBackend writes a minimal Bun API backend in /backend.
func WriteBunBackend() error {
	if err := os.MkdirAll("backend", 0o755); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join("backend", "index.ts"), []byte(templates.BunBackendServer()), 0o644); err != nil {
		return err
	}

	return os.WriteFile(filepath.Join("backend", "package.json"), []byte(templates.BunBackendPackageJSON()), 0o644)
}
