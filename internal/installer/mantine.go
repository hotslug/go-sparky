package installer

import (
	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
)

// InstallMantine installs Mantine dependencies.
func InstallMantine(p plan.Plan) error {
	spin := logger.StartSpinner("Installing Mantine packages")
	if err := addDependencies(p, false,
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
	if err := addDependencies(p, true,
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

// RemoveMantine removes Mantine dependencies and related PostCSS plugins.
func RemoveMantine(p plan.Plan) error {
	spin := logger.StartSpinner("Removing Mantine packages")
	if err := removeDependencies(p, false,
		"@mantine/core",
		"@mantine/hooks",
		"@mantine/form",
		"@mantine/dates",
		"dayjs",
		"@mantine/charts",
		"recharts",
		"@mantine/notifications",
		"@mantine/code-highlight",
		"@mantine/tiptap",
		"@tiptap/pm",
		"@tiptap/react",
		"@tiptap/extension-link",
		"@tiptap/starter-kit",
		"@mantine/dropzone",
		"@mantine/carousel",
		"embla-carousel",
		"embla-carousel-react",
		"@mantine/spotlight",
		"@mantine/modals",
		"@mantine/nprogress",
	); err != nil {
		spin("Failed to remove Mantine")
		return err
	}
	spin("Removed Mantine packages")

	spin = logger.StartSpinner("Removing Mantine PostCSS plugins")
	if err := removeDependencies(p, true,
		"postcss",
		"postcss-preset-mantine",
		"postcss-simple-vars",
	); err != nil {
		spin("Failed to remove Mantine PostCSS plugins")
		return err
	}
	spin("Removed Mantine PostCSS plugins")
	return nil
}
