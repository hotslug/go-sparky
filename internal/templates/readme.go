package templates

import (
	"fmt"
	"strings"

	"github.com/hotslug/go-sparky/internal/plan"
)

// Readme builds a README tailored to the selected features.
func Readme(p plan.Plan) string {
	var b strings.Builder

	title := p.Name
	if title == "" {
		title = "New App"
	}

	fmt.Fprintf(&b, "# %s\n\n", title)
	b.WriteString("Scaffolded with Go Sparky (Vite + React + TypeScript).\n\n")

	b.WriteString("## What's inside\n")
	features := []string{
		"Vite + React + TypeScript",
		"Framer Motion",
	}
	if p.Tailwind {
		features = append(features, "Tailwind CSS")
	}
	if p.Mantine {
		features = append(features, "Mantine UI (with Mantine PostCSS preset)")
	}
	if p.Zustand {
		features = append(features, "Zustand state store (demo slice in src/stores/useSparkyStore.ts)")
	}
	if p.ReactQuery {
		features = append(features, "TanStack Query + Devtools")
	}
	if p.Eslint {
		features = append(features, "ESLint (React, TypeScript, a11y, import order, Prettier)")
	}
	if p.Prettier {
		features = append(features, "Prettier (Tailwind + import sort plugins)")
	}
	if p.Husky {
		features = append(features, "Husky + lint-staged pre-commit")
	}
	if p.Storybook {
		features = append(features, "Storybook (Vite + React config; starter story in src/stories)")
	}
	if p.Docker {
		features = append(features, "Dockerfile + docker-compose (dev/prod)")
	}
	if p.Vercel {
		features = append(features, "Vercel static deploy config")
	}
	if p.Netlify {
		features = append(features, "Netlify SPA deploy config")
	}

	for _, f := range features {
		fmt.Fprintf(&b, "- %s\n", f)
	}
	b.WriteString("\n")

	b.WriteString("## Quickstart\n")
	b.WriteString("```sh\n")
	b.WriteString("pnpm dev\n")
	b.WriteString("```\n\n")
	b.WriteString("Then open http://localhost:5173\n\n")

	b.WriteString("## Scripts\n")
	b.WriteString("- `pnpm dev` – start dev server\n")
	b.WriteString("- `pnpm build` – production build\n")
	b.WriteString("- `pnpm test` – run unit tests\n")
	if p.Eslint {
		b.WriteString("- `pnpm lint` – run ESLint\n")
	}
	if p.Prettier {
		b.WriteString("- `pnpm format` (optional) – run Prettier\n")
	}
	if p.Storybook {
		b.WriteString("- `pnpm storybook dev -p 6006` – run Storybook\n")
	}
	b.WriteString("\n")

	if p.Mantine {
		b.WriteString("## Mantine\n")
		b.WriteString("- Styles imported in `src/App.tsx`\n")
		b.WriteString("- MantineProvider set up in `src/main.tsx`\n\n")
	}

	if p.Zustand {
		b.WriteString("## Zustand\n")
		b.WriteString("- Demo slice: `src/stores/useSparkyStore.ts`\n")
		b.WriteString("- Replace or split slices to match your app state\n\n")
	}

	if p.Tailwind {
		b.WriteString("## Tailwind\n")
		b.WriteString("- Configured via `@tailwindcss/vite`\n")
		b.WriteString("- Styles in `src/index.css`\n\n")
	}

	if p.Docker {
		b.WriteString("## Docker\n")
		b.WriteString("- Dev: `docker compose up dev` (http://localhost:5173)\n")
		b.WriteString("- Prod: `docker compose up prod` (http://localhost:4173)\n\n")
	}

	if p.Vercel || p.Netlify {
		b.WriteString("## Deploy\n")
		if p.Vercel {
			b.WriteString("- Vercel: `vercel --prod` (uses `vercel.json`)\n")
		}
		if p.Netlify {
			b.WriteString("- Netlify: `netlify deploy --prod` (uses `netlify.toml`)\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Editing\n")
	b.WriteString("- Main app entry: `src/App.tsx`\n")
	b.WriteString("- Providers/root: `src/main.tsx`\n")

	return b.String()
}
