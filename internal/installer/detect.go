package installer

import (
	"bytes"
	"os"
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
