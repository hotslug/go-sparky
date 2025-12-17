package runner

import (
	"fmt"
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

// RunEnv executes a command with additional environment variables.
func RunEnv(cmd string, env map[string]string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	command.Env = os.Environ()
	for k, v := range env {
		command.Env = append(command.Env, fmt.Sprintf("%s=%s", k, v))
	}

	return command.Run()
}
