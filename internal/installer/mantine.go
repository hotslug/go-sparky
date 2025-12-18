package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/runner"
)

// InstallMantine installs Mantine dependencies.
func InstallMantine() error {
	spin := logger.StartSpinner("Installing Mantine packages")
	if err := runner.RunQuiet(
		"pnpm",
		"install",
		"@mantine/core@latest",
		"@mantine/hooks@latest",
		"@mantine/form@latest",
		"@mantine/dates@latest",
		"dayjs@latest",
		"@mantine/charts@latest",
		"recharts@latest",
		"@mantine/notifications@latest",
		"@mantine/code-highlight@latest",
		"@mantine/tiptap@latest",
		"@tiptap/pm@latest",
		"@tiptap/react@latest",
		"@tiptap/extension-link@latest",
		"@tiptap/starter-kit@latest",
		"@mantine/dropzone@latest",
		"@mantine/carousel@latest",
		"embla-carousel@^8.5.2",
		"embla-carousel-react@^8.5.2",
		"@mantine/spotlight@latest",
		"@mantine/modals@latest",
		"@mantine/nprogress@latest",
	); err != nil {
		spin("Failed to install Mantine")
		return err
	}
	spin("Installed Mantine packages")

	spin = logger.StartSpinner("Installing Mantine PostCSS plugins")
	if err := runner.RunQuiet(
		"pnpm",
		"install",
		"-D",
		"postcss@latest",
		"postcss-preset-mantine@latest",
		"postcss-simple-vars@latest",
	); err != nil {
		spin("Failed to install Mantine PostCSS plugins")
		return err
	}
	spin("Installed Mantine PostCSS plugins")
	return nil
}
