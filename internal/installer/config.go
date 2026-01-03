package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/plan"
)

const postcssConfigContent = `module.exports = {
  plugins: {
    'postcss-preset-mantine': {},
    'postcss-simple-vars': {
      variables: {
        'mantine-breakpoint-xs': '36em',
        'mantine-breakpoint-sm': '48em',
        'mantine-breakpoint-md': '62em',
        'mantine-breakpoint-lg': '75em',
        'mantine-breakpoint-xl': '88em',
      },
    },
  },
};
`

// WriteViteConfig writes vite.config.ts, optionally including the Tailwind plugin.
func WriteViteConfig(includeTailwind bool) error {
	content := `import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
`

	if includeTailwind {
		content += `import tailwindcss from "@tailwindcss/vite";
`
	}

	content += `import path from "node:path";
import { fileURLToPath } from "node:url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default defineConfig({
  plugins: [react()`

	if includeTailwind {
		content += ", tailwindcss()"
	}

	content += `],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
});
`

	return os.WriteFile("vite.config.ts", []byte(content), 0o644)
}

// WriteConfigFiles writes bundler-specific config files.
func WriteConfigFiles(p plan.Plan) error {
	if p.IsVite() {
		return WriteViteConfig(p.Tailwind)
	}

	return nil
}

// WritePostCSSConfig writes postcss.config.cjs with a lightweight Mantine preset.
func WritePostCSSConfig() error {
	return os.WriteFile("postcss.config.cjs", []byte(postcssConfigContent), 0o644)
}

// DeletePostCSSConfigIfOwned deletes postcss.config.cjs when it matches our generated content.
func DeletePostCSSConfigIfOwned() error {
	data, err := os.ReadFile("postcss.config.cjs")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if string(data) != postcssConfigContent {
		return nil
	}

	return os.Remove("postcss.config.cjs")
}
