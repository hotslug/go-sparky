package logger

import (
	"fmt"
)

// Step prints a human-friendly progress message.
func Step(msg string) {
	fmt.Printf("• %s\n", msg)
}

// Info prints informational text.
func Info(msg string) {
	fmt.Println(msg)
}

// Success prints a success message.
func Success(msg string) {
	fmt.Printf("✅ %s\n", msg)
}

// Error prints an error message.
func Error(msg string) {
	fmt.Printf("✖️ %s\n", msg)
}
