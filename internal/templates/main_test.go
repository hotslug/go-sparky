package templates

import (
	"strings"
	"testing"

	"github.com/hotslug/go-sparky/internal/plan"
)

func TestMainTemplateWithProviders(t *testing.T) {
	p := plan.Plan{Mantine: true, ReactQuery: true}
	content := MainTemplate(p)

	checkIncludes(t, content, "import { MantineProvider } from '@mantine/core';")
	checkIncludes(t, content, "import { QueryClient, QueryClientProvider } from '@tanstack/react-query';")

	queryIdx := strings.Index(content, "<QueryClientProvider")
	mantineIdx := strings.Index(content, "<MantineProvider>")
	if queryIdx == -1 || mantineIdx == -1 || queryIdx > mantineIdx {
		t.Fatalf("expected QueryClientProvider to wrap MantineProvider")
	}

	devtoolsIdx := strings.Index(content, "<ReactQueryDevtools")
	if devtoolsIdx == -1 || devtoolsIdx < mantineIdx {
		t.Fatalf("expected ReactQueryDevtools after Mantine provider block")
	}
}

func TestMainTemplateMantineOnly(t *testing.T) {
	p := plan.Plan{Mantine: true}
	content := MainTemplate(p)

	if strings.Contains(content, "QueryClientProvider") {
		t.Fatalf("did not expect QueryClientProvider")
	}
	checkIncludes(t, content, "<MantineProvider>")
}

func TestMainTemplateReactQueryOnly(t *testing.T) {
	p := plan.Plan{ReactQuery: true}
	content := MainTemplate(p)

	if strings.Contains(content, "MantineProvider") {
		t.Fatalf("did not expect MantineProvider")
	}
	checkIncludes(t, content, "<QueryClientProvider client={queryClient}>")
	checkIncludes(t, content, "<ReactQueryDevtools initialIsOpen={false} />")
}

func checkIncludes(t *testing.T, content, needle string) {
	t.Helper()
	if !strings.Contains(content, needle) {
		t.Fatalf("expected content to include %q", needle)
	}
}
