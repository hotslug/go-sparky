package templates

import (
	"testing"

	"github.com/hotslug/go-sparky/internal/plan"
)

func TestAppTemplateSelection(t *testing.T) {
	t.Run("styled mantine", func(t *testing.T) {
		p := plan.Plan{Mantine: true, StyledApp: true, Zustand: true}
		got := AppTemplate(p)
		if got != styledMantineApp {
			t.Fatalf("expected styled Mantine template")
		}
	})

	t.Run("basic mantine", func(t *testing.T) {
		p := plan.Plan{Mantine: true}
		got := AppTemplate(p)
		if got != basicApp {
			t.Fatalf("expected basic template when Mantine without styled")
		}
	})

	t.Run("zustand default", func(t *testing.T) {
		p := plan.Plan{Zustand: true}
		got := AppTemplate(p)
		if got != zustandApp {
			t.Fatalf("expected zustand template when zustand enabled")
		}
	})

	t.Run("basic react", func(t *testing.T) {
		p := plan.Plan{}
		got := AppTemplate(p)
		if got != basicApp {
			t.Fatalf("expected basic React template")
		}
	})
}
