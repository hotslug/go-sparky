package installer

import (
	"os"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/templates"
)

// WriteAppFiles updates the scaffolded source files with our templates.
func WriteAppFiles(p plan.Plan) error {
	if err := os.WriteFile(filepath.Join("src", "App.tsx"), []byte(templates.AppTemplate(p)), 0o644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join("src", "main.tsx"), []byte(templates.MainTemplate(p)), 0o644); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join("src", "assets"), 0o755); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join("src", "assets", "sparky.png"), sparkyImage, 0o644); err != nil {
		return err
	}

	css := baseIndexCSS
	if p.Tailwind {
		css = tailwindIndexCSS
	}

	return os.WriteFile(filepath.Join("src", "index.css"), []byte(css), 0o644)
}

const baseIndexCSS = `@import url('https://fonts.googleapis.com/css2?family=Fredoka:wght@400;600;700&display=swap');

:root {
  font-family: 'Inter', system-ui, Avenir, Helvetica, Arial, sans-serif;
  line-height: 1.5;
  font-weight: 400;
  color: #0f172a;
  background-color: #f8fafc;
}

body {
  margin: 0;
  min-height: 100vh;
}

#root {
  min-height: 100vh;
}

.font-sparky {
  font-family: 'Fredoka', ui-rounded, system-ui, sans-serif;
}
`

const tailwindIndexCSS = `@import url('https://fonts.googleapis.com/css2?family=Fredoka:wght@400;600;700&display=swap');
@import "tailwindcss";

:root {
  color: #e2e8f0;
  background-color: #0f172a;
  font-family: 'Inter', system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
}

body {
  margin: 0;
  min-height: 100vh;
}

#root {
  min-height: 100vh;
}

.font-sparky {
  font-family: 'Fredoka', ui-rounded, system-ui, sans-serif;
}
`
