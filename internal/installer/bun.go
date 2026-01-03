package installer

import (
	"fmt"
	"os"
	"strings"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/hotslug/go-sparky/internal/plan"
	"github.com/hotslug/go-sparky/internal/runner"
)

// ScaffoldBunProject runs bun init --react --yes inside the target directory.
func ScaffoldBunProject() error {
	spin := logger.StartSpinner("Scaffolding with Bun (React + TypeScript)")
	if err := runner.RunQuiet("bun", "init", "--react", "--yes"); err != nil {
		spin("Failed to scaffold project")
		return err
	}
	spin("Scaffolded Bun project")
	return nil
}

// CleanupBunScaffold removes Bun template files we'll replace.
// Version: Bun 1.3.5 (update this comment when Bun updates).
func CleanupBunScaffold() error {
	filesToRemove := []string{
		"CLAUDE.md",
		".cursor",
		"README.md",
		"src/App.tsx",
		"src/frontend.tsx",
		"src/APITester.tsx",
		"src/react.svg",
		"src/logo.svg",
	}

	for _, file := range filesToRemove {
		_ = os.RemoveAll(file)
	}

	return nil
}

// WriteBunConfig updates bunfig.toml with plugins needed by the plan.
func WriteBunConfig(p plan.Plan) error {
	if !p.Tailwind {
		return nil
	}

	const pluginName = "bun-plugin-tailwind"

	data, err := os.ReadFile("bunfig.toml")
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if strings.Contains(string(data), pluginName) {
		return nil
	}

	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !strings.HasPrefix(trimmed, "plugins") {
			continue
		}

		open := strings.Index(line, "[")
		close := strings.Index(line, "]")
		if open == -1 || close == -1 || close < open {
			return fmt.Errorf("unsupported bunfig.toml plugins format; add %q manually", pluginName)
		}

		contents := strings.TrimSpace(line[open+1 : close])
		if contents == "" {
			lines[i] = line[:open+1] + fmt.Sprintf("%q", pluginName) + line[close:]
		} else {
			lines[i] = line[:close] + ", " + fmt.Sprintf("%q", pluginName) + line[close:]
		}

		return os.WriteFile("bunfig.toml", []byte(strings.Join(lines, "\n")), 0o644)
	}

	appendLine := "plugins = [" + fmt.Sprintf("%q", pluginName) + "]"
	if len(data) == 0 {
		return os.WriteFile("bunfig.toml", []byte(appendLine+"\n"), 0o644)
	}

	return os.WriteFile("bunfig.toml", []byte(strings.TrimRight(string(data), "\n")+"\n"+appendLine+"\n"), 0o644)
}
