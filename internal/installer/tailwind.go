package installer

import (
	"os"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallTailwind installs Tailwind dependencies and writes related configs.
func InstallTailwind() error {
	spin := logger.StartSpinner("Installing Tailwind CSS")
	if err := runner.RunQuiet("pnpm", "install", "-D", "tailwindcss@latest", "@tailwindcss/vite@latest", "@tailwindcss/postcss@latest"); err != nil {
		spin("Failed to install Tailwind CSS")
		return err
	}
	spin("Installed Tailwind CSS")

	return os.WriteFile("tailwind.config.ts", []byte(tailwindConfig), 0o644)
}

const tailwindConfig = `import type { Config } from 'tailwindcss';

export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      colors: {
        brand: '#5f3dc4',
      },
      borderRadius: {
        base: '12px',
      },
    },
  },
  plugins: [],
} satisfies Config;
`
