package installer

import (
	"os"
)

// WriteViteConfig writes vite.config.ts, optionally including the Tailwind plugin.
func WriteViteConfig(includeTailwind bool) error {
	content := `import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
`

	if includeTailwind {
		content += "import tailwindcss from '@tailwindcss/vite';\n\n"
		content += "export default defineConfig({\n"
		content += "  plugins: [react(), tailwindcss()],\n"
		content += "});\n"
	} else {
		content += "\nexport default defineConfig({\n"
		content += "  plugins: [react()],\n"
		content += "});\n"
	}

	return os.WriteFile("vite.config.ts", []byte(content), 0o644)
}

// WritePostCSSConfig writes postcss.config.cjs with a lightweight Mantine preset.
func WritePostCSSConfig(includeTailwind bool) error {
	plugins := "  plugins: [\n    mantinePreset,\n"
	if includeTailwind {
		plugins += "    require('tailwindcss'),\n"
	}
	plugins += "  ],\n"

	content := `const mantinePreset = () => ({
  postcssPlugin: 'mantine-preset',
  Once(root) {
    root.prepend(':root { --mantine-accent: #339af0; --mantine-radius: 12px; }');
  },
});
mantinePreset.postcss = true;

module.exports = {
`
	content += plugins
	content += "};\n"

	return os.WriteFile("postcss.config.cjs", []byte(content), 0o644)
}
