package installer

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/hotslug/go-sparky/internal/logger"
    "github.com/hotslug/go-sparky/internal/runner"
)

// CreateInitialCommitIfMissing commits the current workspace if no commits exist yet.
func CreateInitialCommitIfMissing(message string) error {
	if _, err := os.Stat(".git"); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	// If a commit already exists, skip.
	if err := exec.Command("git", "rev-parse", "--verify", "HEAD").Run(); err == nil {
		return nil
	}

	logger.Step("Creating initial git commit")

	if err := runner.RunQuiet("git", "add", "-A"); err != nil {
		return err
	}

	// If nothing was staged, exit quietly.
	if err := exec.Command("git", "diff", "--cached", "--quiet").Run(); err == nil {
		return nil
	} else if _, ok := err.(*exec.ExitError); !ok {
		return err
	}

	if err := runner.RunQuiet("git", "commit", "-m", message); err != nil {
		return fmt.Errorf("Initial git commit failed (check hook output above or rerun with --no-husky): %w", err)
	}

	return nil
}
