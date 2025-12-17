package runner

import (
	"bytes"
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

// RunQuiet executes a command silently, only showing output if there's an error.
func RunQuiet(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)

	var outBuf, errBuf bytes.Buffer
	command.Stdout = &outBuf
	command.Stderr = &errBuf

	err := command.Run()
	if err != nil {
		// Print captured output on error
		if outBuf.Len() > 0 {
			fmt.Print(outBuf.String())
		}
		if errBuf.Len() > 0 {
			fmt.Fprint(os.Stderr, errBuf.String())
		}
		return err
	}

	return nil
}

// RunQuietEnv executes a command with environment variables silently, only showing output on error.
func RunQuietEnv(cmd string, env map[string]string, args ...string) error {
	command := exec.Command(cmd, args...)

	var outBuf, errBuf bytes.Buffer
	command.Stdout = &outBuf
	command.Stderr = &errBuf
	command.Stdin = nil

	command.Env = os.Environ()
	for k, v := range env {
		command.Env = append(command.Env, fmt.Sprintf("%s=%s", k, v))
	}

	err := command.Run()
	if err != nil {
		// Print captured output on error
		if outBuf.Len() > 0 {
			fmt.Print(outBuf.String())
		}
		if errBuf.Len() > 0 {
			fmt.Fprint(os.Stderr, errBuf.String())
		}
		return err
	}

	return nil
}

// RunEnv executes a command with additional environment variables.
func RunEnv(cmd string, env map[string]string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	// Don't attach stdin to prevent interactive prompts when CI=1
	command.Stdin = nil

	command.Env = os.Environ()
	for k, v := range env {
		command.Env = append(command.Env, fmt.Sprintf("%s=%s", k, v))
	}

	return command.Run()
}
