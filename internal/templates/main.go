package templates

import (
	"strings"

	"github.com/hotslug/go-sparky/internal/plan"
)

// MainTemplate builds main.tsx with conditional providers.
func MainTemplate(p plan.Plan) string {
	var imports []string

	imports = append(imports,
		"import React from 'react';",
		"import ReactDOM from 'react-dom/client';",
		"import App from './App';",
		"import './index.css';",
	)

	if p.Mantine {
		imports = append(imports, "import { MantineProvider } from '@mantine/core';")
	}

	if p.ReactQuery {
		imports = append(imports,
			"import { QueryClient, QueryClientProvider } from '@tanstack/react-query';",
			"import { ReactQueryDevtools } from '@tanstack/react-query-devtools';",
		)
	}

	var b strings.Builder

	for _, line := range imports {
		b.WriteString(line)
		b.WriteString("\n")
	}

	b.WriteString("\n")

	if p.ReactQuery {
		b.WriteString("const queryClient = new QueryClient();\n\n")
	}

	b.WriteString("const rootElement = document.getElementById('root');\n")
	b.WriteString("if (!rootElement) throw new Error('Root element not found');\n")
	b.WriteString("const root = ReactDOM.createRoot(rootElement);\n\n")

	switch {
	case p.ReactQuery && p.Mantine:
		b.WriteString("root.render(\n")
		b.WriteString("  <React.StrictMode>\n")
		b.WriteString("    <QueryClientProvider client={queryClient}>\n")
		b.WriteString("      <MantineProvider>\n")
		b.WriteString("        <App />\n")
		b.WriteString("      </MantineProvider>\n")
		b.WriteString("      <ReactQueryDevtools initialIsOpen={false} />\n")
		b.WriteString("    </QueryClientProvider>\n")
		b.WriteString("  </React.StrictMode>\n")
		b.WriteString(");\n")
	case p.ReactQuery:
		b.WriteString("root.render(\n")
		b.WriteString("  <React.StrictMode>\n")
		b.WriteString("    <QueryClientProvider client={queryClient}>\n")
		b.WriteString("      <App />\n")
		b.WriteString("      <ReactQueryDevtools initialIsOpen={false} />\n")
		b.WriteString("    </QueryClientProvider>\n")
		b.WriteString("  </React.StrictMode>\n")
		b.WriteString(");\n")
	case p.Mantine:
		b.WriteString("root.render(\n")
		b.WriteString("  <React.StrictMode>\n")
		b.WriteString("    <MantineProvider>\n")
		b.WriteString("      <App />\n")
		b.WriteString("    </MantineProvider>\n")
		b.WriteString("  </React.StrictMode>\n")
		b.WriteString(");\n")
	default:
		b.WriteString("root.render(\n")
		b.WriteString("  <React.StrictMode>\n")
		b.WriteString("    <App />\n")
		b.WriteString("  </React.StrictMode>\n")
		b.WriteString(");\n")
	}

	return b.String()
}
