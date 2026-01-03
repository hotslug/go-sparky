package installer

import (
	"bytes"
	"fmt"
	"os"

	"github.com/hotslug/go-sparky/internal/plan"
)

// HasReactQueryDependency reports whether package.json lists @tanstack/react-query.
// Used to preserve React Query wiring when updating templates on existing projects.
func HasReactQueryDependency() bool {
	data, err := os.ReadFile("package.json")
	if err != nil {
		return false
	}

	return bytes.Contains(data, []byte("@tanstack/react-query"))
}

// HasMantineDependency reports whether package.json lists @mantine/core.
func HasMantineDependency() bool {
	data, err := os.ReadFile("package.json")
	if err != nil {
		return false
	}

	return bytes.Contains(data, []byte("@mantine/core"))
}

// HasZustandDependency reports whether package.json lists zustand.
func HasZustandDependency() bool {
	data, err := os.ReadFile("package.json")
	if err != nil {
		return false
	}

	return bytes.Contains(data, []byte("\"zustand\""))
}

// HasTailwind reports whether tailwindcss appears in package.json or a tailwind config exists.
func HasTailwind() bool {
	if hasTailwindPackage() {
		return true
	}

	configs := []string{
		"tailwind.config.js",
		"tailwind.config.cjs",
		"tailwind.config.mjs",
		"tailwind.config.ts",
	}

	for _, cfg := range configs {
		if _, err := os.Stat(cfg); err == nil {
			return true
		}
	}

	return false
}

func hasTailwindPackage() bool {
	data, err := os.ReadFile("package.json")
	if err != nil {
		return false
	}

	return bytes.Contains(data, []byte("tailwindcss"))
}

// DetectBundler returns the bundler type, preferring explicit markers.
// Precedence: Vite config > Bun markers > error (no bundler detected).
func DetectBundler() (plan.BundlerType, error) {
	hasVite := HasViteConfig()
	hasBun := HasBunProject()

	if hasVite {
		return plan.BundlerVite, nil
	}

	if hasBun {
		return plan.BundlerBun, nil
	}

	return "", fmt.Errorf("no bundler detected: run this command from a go-sparky project root")
}

// HasBunProject reports whether Bun markers exist in the project.
func HasBunProject() bool {
	bunMarkers := []string{
		"bunfig.toml",
		"bun.lock",
		"bun-env.d.ts",
	}

	for _, marker := range bunMarkers {
		if _, err := os.Stat(marker); err == nil {
			return true
		}
	}

	return false
}

// HasViteConfig reports whether a Vite config file exists.
func HasViteConfig() bool {
	viteConfigs := []string{
		"vite.config.ts",
		"vite.config.js",
		"vite.config.mjs",
	}

	for _, cfg := range viteConfigs {
		if _, err := os.Stat(cfg); err == nil {
			return true
		}
	}

	return false
}
