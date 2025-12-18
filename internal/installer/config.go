package installer

import (
	"os"
)

// WriteViteConfig writes vite.config.ts, optionally including the Tailwind plugin.
func WriteViteConfig(includeTailwind bool) error {
	content := `import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
`

	if includeTailwind {
		content += `import tailwindcss from "@tailwindcss/vite";
`
	}

	content += `import path from "path";

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

// WritePostCSSConfig writes postcss.config.cjs with a lightweight Mantine preset.
func WritePostCSSConfig() error {
	content := `module.exports = {
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

	return os.WriteFile("postcss.config.cjs", []byte(content), 0o644)
}
