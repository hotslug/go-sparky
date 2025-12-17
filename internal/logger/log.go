package logger

import "fmt"

var verbose bool

// SetVerbose toggles verbose output helpers (e.g., spinner).
func SetVerbose(v bool) {
	verbose = v
}

// Step prints a human-friendly progress message.
func Step(msg string) {
	fmt.Printf("\033[1;36m•\033[0m %s\n", msg)
}

// Info prints informational text.
func Info(msg string) {
	fmt.Println(msg)
}

// Success prints a success message.
func Success(msg string) {
	fmt.Printf("\033[32m✅ %s\033[0m\n", msg)
}

// Error prints an error message.
func Error(msg string) {
	fmt.Printf("\033[31m✖️  %s\033[0m\n", msg)
}

// Warning prints a warning message.
func Warning(msg string) {
	fmt.Printf("\033[33m⚠️  %s\033[0m\n", msg)
}
