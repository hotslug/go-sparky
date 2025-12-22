package installer

import (
	"os"
	"path/filepath"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallStorybook installs Storybook dev dependencies for Vite + React.
func InstallStorybook() error {
	spin := logger.StartSpinner("Installing Storybook")
	if err := runner.RunQuiet(
		"pnpm",
		"install",
		"-D",
		"storybook@latest",
		"@storybook/react-vite@latest",
		"@storybook/addon-essentials@latest",
		"@storybook/addon-interactions@latest",
		"@storybook/blocks@latest",
		"@storybook/test@latest",
	); err != nil {
		spin("Failed to install Storybook")
		return err
	}
	spin("Installed Storybook")

	return nil
}

// WriteStorybookConfig writes .storybook config files and a starter story.
// includeIndexCSS toggles importing src/index.css in preview.ts when it exists.
func WriteStorybookConfig(includeIndexCSS bool) error {
	if err := os.MkdirAll(".storybook", 0o755); err != nil {
		return err
	}

	mainContent := `import type { StorybookConfig } from "@storybook/react-vite";

const config: StorybookConfig = {
  stories: ["../src/**/*.mdx", "../src/**/*.stories.@(js|jsx|ts|tsx)"],
  addons: [
    "@storybook/addon-essentials",
    "@storybook/addon-interactions",
    "@storybook/blocks",
  ],
  framework: {
    name: "@storybook/react-vite",
    options: {},
  },
  docs: {
    autodocs: "tag",
  },
};

export default config;
`

	previewContent := `import type { Preview } from "@storybook/react";
`
	if includeIndexCSS {
		previewContent += `import "../src/index.css";
`
	}

	previewContent += `
const preview: Preview = {
  parameters: {
    actions: { argTypesRegex: "^on[A-Z].*" },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
};

export default preview;
`

	if err := os.WriteFile(filepath.Join(".storybook", "main.ts"), []byte(mainContent), 0o644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(".storybook", "preview.ts"), []byte(previewContent), 0o644); err != nil {
		return err
	}

	storyDir := filepath.Join("src", "stories")
	if err := os.MkdirAll(storyDir, 0o755); err != nil {
		return err
	}

	storyPath := filepath.Join(storyDir, "SparkyCard.stories.tsx")
	if _, err := os.Stat(storyPath); err == nil {
		return nil
	} else if err != nil && !os.IsNotExist(err) {
		return err
	}

	storyContent := `import type { Meta, StoryObj } from "@storybook/react";
import sparky from "../assets/sparky.png";

type SparkyCardProps = {
  title: string;
  tagline: string;
  cta: string;
};

function SparkyCard({ title, tagline, cta }: SparkyCardProps) {
  return (
    <div className="max-w-lg rounded-2xl border border-slate-200/30 bg-white/5 p-6 shadow-lg backdrop-blur">
      <div className="flex items-center gap-4">
        <img
          src={sparky}
          alt="Go Sparky mascot"
          className="h-28 w-28 rounded-xl border border-slate-200/40 bg-slate-900/40 object-cover"
        />
        <div className="space-y-2">
          <h2 className="text-2xl font-semibold text-slate-50">{title}</h2>
          <p className="text-slate-200">{tagline}</p>
          <button className="inline-flex items-center gap-2 rounded-lg bg-gradient-to-r from-blue-500 to-purple-500 px-3 py-2 text-sm font-semibold text-white shadow-sm transition hover:shadow-md hover:shadow-blue-500/30">
            {cta}
          </button>
        </div>
      </div>
    </div>
  );
}

const meta: Meta<typeof SparkyCard> = {
  title: "Sparky/Card",
  component: SparkyCard,
  args: {
    title: "Go-Sparky",
    tagline: "Vite + React + TypeScript + Tailwind in one CLI",
    cta: "Get started",
  },
};

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {};

export const BoldCTA: Story = {
  args: {
    cta: "Launch dev server",
  },
};
`

	return os.WriteFile(storyPath, []byte(storyContent), 0o644)
}

// HasStorybookConfig checks if .storybook already exists.
func HasStorybookConfig() bool {
	_, err := os.Stat(".storybook")
	return err == nil
}
