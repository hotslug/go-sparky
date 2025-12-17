package installer

import "os"

// WriteVercelConfig writes a minimal static build config for Vercel.
func WriteVercelConfig() error {
	return os.WriteFile("vercel.json", []byte(vercelConfig), 0o644)
}

// WriteNetlifyConfig writes a basic Netlify deploy config.
func WriteNetlifyConfig() error {
	return os.WriteFile("netlify.toml", []byte(netlifyConfig), 0o644)
}

const vercelConfig = `{
  "builds": [
    {
      "src": "package.json",
      "use": "@vercel/static-build",
      "config": { "distDir": "dist" }
    }
  ],
  "devCommand": "pnpm dev",
  "buildCommand": "pnpm build",
  "outputDirectory": "dist"
}
`

const netlifyConfig = `[build]
command = "pnpm build"
publish = "dist"

[[redirects]]
from = "/*"
to = "/index.html"
status = 200
`
