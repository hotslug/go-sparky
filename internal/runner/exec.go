package runner

import (
	"os"
	"os/exec"
)

// Run executes a command and streams stdout/stderr directly.
func Run(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	return command.Run()
}
