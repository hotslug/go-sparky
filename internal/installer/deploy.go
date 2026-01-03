package installer

import (
	"fmt"
	"os"

	"github.com/hotslug/go-sparky/internal/plan"
)

// WriteVercelConfig writes a minimal static build config for Vercel.
func WriteVercelConfig(p plan.Plan) error {
	return os.WriteFile("vercel.json", []byte(VercelConfig(p)), 0o644)
}

// WriteNetlifyConfig writes a basic Netlify deploy config.
func WriteNetlifyConfig(p plan.Plan) error {
	return os.WriteFile("netlify.toml", []byte(NetlifyConfig(p)), 0o644)
}

// VercelConfig returns a static build config for the chosen bundler.
func VercelConfig(p plan.Plan) string {
	cmd := fmt.Sprintf("%s dev", p.PackageManager())
	build := fmt.Sprintf("%s build", p.PackageManager())
	if p.IsBun() {
		cmd = "bun run dev"
		build = "bun run build"
	}

	return `{
  "builds": [
    {
      "src": "package.json",
      "use": "@vercel/static-build",
      "config": { "distDir": "dist" }
    }
  ],
  "devCommand": "` + cmd + `",
  "buildCommand": "` + build + `",
  "outputDirectory": "dist"
}
`
}

// NetlifyConfig returns a Netlify config for the chosen bundler.
func NetlifyConfig(p plan.Plan) string {
	build := fmt.Sprintf("%s build", p.PackageManager())
	if p.IsBun() {
		build = "bun run build"
	}

	return `[build]
command = "` + build + `"
publish = "dist"

[[redirects]]
from = "/*"
to = "/index.html"
status = 200
`
}
