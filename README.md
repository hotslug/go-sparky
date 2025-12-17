# go-sparky

CLI scaffolder for React + Vite + TypeScript that installs opinionated options (Mantine, Tailwind, TanStack Query, ESLint, Prettier, Husky) and rewrites the starter files with matching templates.

## Prerequisites
- Go 1.21+ (to build/run the CLI)
- Node 18+ with `pnpm` on your PATH (used to generate and install the Vite app)

## Installation
- Latest released binary via Go: `go install github.com/hotslug/go-sparky@latest`
- From source in this repo: `go run . --help` or `go build .` then run `./go-sparky`

## Usage
Scaffold a new app into a fresh directory:

```sh
go-sparky new my-app --mantine --tailwind --react-query --eslint --prettier --husky --styled
```

Flags:
- `--mantine` – add Mantine UI and wrap the app in `MantineProvider`
- `--tailwind` – install Tailwind and write `tailwind.config.ts`/`index.css`
- `--react-query` – add TanStack Query + devtools and wire providers
- `--eslint` – install ESLint with React/DOM presets
- `--prettier` – install Prettier with Vite-friendly defaults
- `--husky` – set up Husky + lint-staged
- `--styled` – use the styled Mantine landing page template (requires `--mantine`)
- `--docker` – add Dockerfile + docker-compose.yml (dev + prod)
- `--vercel` – add `vercel.json` for static deploys
- `--netlify` – add `netlify.toml` with SPA redirect

After scaffolding:
- `cd my-app`
- `pnpm dev`
- edit `src/App.tsx` to start building ⚡

Docker quickstart:
- `docker compose up dev` for hot-reload dev server on `localhost:5173`
- `docker compose up prod` to build & serve the production bundle at `localhost:4173`

## Development
- Run tests: `go test ./...`
- The ASCII banner lives at `assets/ascii/sparky.txt`; CLI entrypoint is `cmd/root.go`.

## License
MIT © Victor Cintron
