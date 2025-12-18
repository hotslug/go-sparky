<img src="assets/sparky.png" alt="Go Sparky mascot" width="220" align="left" />

<h1>go-sparky</h1>

<br clear="left" />

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
- `--mantine` – add Mantine UI and wrap the app in `MantineProvider` (enables PostCSS preset)
- `--no-tailwind` – skip Tailwind (default installs)
- `--no-react-query` – skip TanStack Query (default installs)
- `--no-eslint` – skip ESLint (default installs)
- `--no-prettier` – skip Prettier (default installs)
- `--no-husky` – skip Husky + lint-staged (default installs)
- `--styled` – use the styled Mantine landing page template (requires `--mantine`)
- `--docker` – add Dockerfile + docker-compose.yml (dev + prod)
- `--vercel` – add `vercel.json` for static deploys
- `--netlify` – add `netlify.toml` with SPA redirect

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
