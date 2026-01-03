package installer

import "github.com/hotslug/go-sparky/internal/plan"

// RemoveReactQuery removes TanStack Query dependencies.
func RemoveReactQuery(p plan.Plan) error {
	return removeDependencies(p, false, "@tanstack/react-query", "@tanstack/react-query-devtools")
}
