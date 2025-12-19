package installer

import "github.com/hotslug/go-sparky/internal/runner"

// RemoveReactQuery removes TanStack Query dependencies.
func RemoveReactQuery() error {
	return runner.RunQuiet("pnpm", "remove", "@tanstack/react-query", "@tanstack/react-query-devtools")
}
