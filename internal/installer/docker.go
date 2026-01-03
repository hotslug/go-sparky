package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/plan"
)

// WriteDockerArtifacts creates Dockerfile and docker-compose.yml for dev/prod flows.
func WriteDockerArtifacts(p plan.Plan) error {
	dockerfileContents := dockerfileViteContents
	dockerComposeContents := dockerComposeViteContents
	if p.IsBun() {
		dockerfileContents = dockerfileBunContents
		dockerComposeContents = dockerComposeBunContents
	}

	if err := os.WriteFile("Dockerfile", []byte(dockerfileContents), 0o644); err != nil {
		return err
	}

	return os.WriteFile("docker-compose.yml", []byte(dockerComposeContents), 0o644)
}

const dockerfileViteContents = `# Build static assets
FROM node:20-alpine AS base
WORKDIR /app
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM base AS deps
COPY package.json pnpm-lock.yaml* ./
RUN pnpm install --frozen-lockfile

FROM base AS build
COPY --from=deps /app/node_modules ./node_modules
COPY . .
RUN pnpm run build

# Serve with nginx
FROM nginx:1.27-alpine AS runner
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
`

const dockerComposeViteContents = `version: "3.9"

services:
  dev:
    image: node:20-alpine
    working_dir: /app
    command: ["pnpm", "run", "dev", "--host", "0.0.0.0", "--port", "5173"]
    ports:
      - "5173:5173"
    volumes:
      - ./:/app
      - /app/node_modules
    environment:
      NODE_ENV: development
    tty: true

  prod:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "4173:80"
`

const dockerfileBunContents = `# Build static assets
FROM oven/bun:1 AS build
WORKDIR /app
COPY package.json bun.lock* ./
RUN bun install --frozen-lockfile
COPY . .
RUN bun run build

# Serve with nginx
FROM nginx:1.27-alpine AS runner
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
`

const dockerComposeBunContents = `version: "3.9"

services:
  dev:
    image: oven/bun:1
    working_dir: /app
    command: ["bun", "run", "dev"]
    ports:
      - "3000:3000"
    volumes:
      - ./:/app
      - /app/node_modules
    environment:
      NODE_ENV: development
    tty: true

  prod:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "4173:80"
`
