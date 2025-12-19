<img src="assets/sparky.png" alt="Go Sparky mascot" width="300" />

# go-sparky

CLI scaffolder for React + Vite + TypeScript that installs an opinionated default stack (Tailwind, TanStack Query, ESLint, Prettier, Husky, Framer Motion) and rewrites the starter files with matching templates. Mantine and deploy configs are optional.

## Prerequisites
- Go 1.21+ (to build/run the CLI)
- Node 20+ with `pnpm` on your PATH (used to generate and install the Vite app)

## Installation
- Latest released binary via Go: `go install github.com/hotslug/go-sparky@latest`
- From source in this repo: `go run . --help` or `go build .` then run `./go-sparky`

## Usage
Scaffold a new app into a fresh directory (defaults to Tailwind, React Query, ESLint, Prettier, Husky, Framer Motion):

```sh
go-sparky new my-app
```

Flags:
- `--mantine` – add Mantine UI and wrap the app in `MantineProvider` (enables PostCSS preset). Uses the default App template unless combined with `--styled`.
- `--no-tailwind` – skip Tailwind (default installs)
- `--no-react-query` – skip TanStack Query (default installs)
- `--no-eslint` – skip ESLint (default installs)
- `--no-prettier` – skip Prettier (default installs)
- `--no-husky` – skip Husky + lint-staged (default installs)
- `--styled` – use the styled Mantine landing page template (requires `--mantine`)
- `--no-framer-motion` – skip Framer Motion (default installs)
- `--docker` – add Dockerfile + docker-compose.yml (dev + prod)
- `--vercel` – add `vercel.json` for static deploys
- `--netlify` – add `netlify.toml` with SPA redirect

Add Mantine to an existing project (leaves `src/App.tsx` untouched):

```sh
go-sparky add mantine
```

This installs Mantine packages, writes `postcss.config.cjs`, and rewires `src/main.tsx` to wrap `MantineProvider` (keeps React Query wiring if its deps are present).
Note: `go-sparky add mantine --styled` is not supported; the styled template is only applied during `go-sparky new --mantine --styled` to avoid overwriting your existing `src/App.tsx`.

Add React Query to an existing project (leaves `src/App.tsx` untouched):

```sh
go-sparky add react-query
```

This installs TanStack Query packages and rewires `src/main.tsx` with `QueryClientProvider` (keeps Mantine wiring if present).

Add deploy artifacts to an existing project:

```sh
go-sparky add docker    # Dockerfile + docker-compose.yml
go-sparky add vercel    # vercel.json
go-sparky add netlify   # netlify.toml
go-sparky add framer-motion  # Framer Motion
go-sparky add shadcn    # shadcn-ui init (interactive)
go-sparky add bulma     # Bulma CSS (+ auto @import in src/index.css)
```

Remove Mantine from an existing project (keeps `src/App.tsx` untouched):

```sh
go-sparky remove mantine
```

This uninstalls Mantine packages, removes the Mantine PostCSS plugins (deletes `postcss.config.cjs` if it matches the generated content), and rewrites `src/main.tsx` to remove `MantineProvider` while keeping React Query wiring if present.

Remove React Query from an existing project (keeps `src/App.tsx` untouched):

```sh
go-sparky remove react-query
```

This uninstalls TanStack Query packages and rewrites `src/main.tsx` to remove `QueryClientProvider` while keeping Mantine wiring if present.

Remove generated deploy artifacts:

```sh
go-sparky remove docker    # removes Dockerfile + docker-compose.yml if unmodified
go-sparky remove vercel    # removes vercel.json if unmodified
go-sparky remove netlify   # removes netlify.toml if unmodified
go-sparky remove framer-motion  # uninstalls framer-motion
go-sparky remove bulma     # uninstalls bulma
```

Add shadcn/ui:

```sh
go-sparky add shadcn
```

This runs the official `shadcn-ui` init on top of Tailwind (interactive prompts). If `components.json` already exists, it skips init and reminds you to add components with `pnpm dlx shadcn-ui@latest add <component>`.

TODO:
- Add opt-in flags for additional CSS frameworks.

After scaffolding:
- dependencies are installed (`pnpm install`)
- dev server starts automatically (`pnpm dev`)
- edit `src/App.tsx` to start building ⚡

Docker quickstart:
- `docker compose up dev` for hot-reload dev server on `localhost:5173`
- `docker compose up prod` to build & serve the production bundle at `localhost:4173`

## Development
- Run tests: `go test ./...`
- The ASCII banner lives at `assets/ascii/sparky.txt`; CLI entrypoint is `cmd/root.go`.

## License
CC0 1.0 Universal
