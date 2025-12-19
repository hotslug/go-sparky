package templates

import (
	"strings"
	"testing"
)

func TestEslintConfigRelaxed_DoesNotReferenceUnicorn(t *testing.T) {
	cfg := EslintConfigRelaxed()
	if strings.Contains(cfg, "unicorn") {
		t.Fatalf("relaxed config should not include unicorn plugin")
	}
}
