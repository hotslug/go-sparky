package templates

import (
	"strings"
	"testing"

	"github.com/hotslug/go-sparky/internal/plan"
)

func TestEslintConfigRelaxed_DoesNotReferenceUnicorn(t *testing.T) {
	cfg := EslintConfigRelaxed(plan.Plan{Bundler: plan.BundlerVite})
	if strings.Contains(cfg, "unicorn") {
		t.Fatalf("relaxed config should not include unicorn plugin")
	}
}
