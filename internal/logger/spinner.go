package logger

import (
	"fmt"
	"time"
)

// StartSpinner prints a simple ASCII spinner and returns a stop function.
// Call the returned function with a final message to end the spinner.
func StartSpinner(msg string) func(finalMsg string) {
	if !verbose {
		return func(string) {}
	}

	stop := make(chan struct{})

	go func() {
		// Chunkier block spinner is easier to see than the thin braille dots.
		frames := []string{"⠋",
			"⠙",
			"⠹",
			"⠸",
			"⠼",
			"⠴",
			"⠦",
			"⠧",
			"⠇",
			"⠏",
		}
		color := "\033[38;5;45m" // bright cyan
		i := 0
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Printf("\r\033[2K%s%s\033[0m %s", color, frames[i%len(frames)], msg)
				i++
				time.Sleep(120 * time.Millisecond)
			}
		}
	}()

	return func(finalMsg string) {
		close(stop)
		fmt.Printf("\r\033[2K\033[32m✓ %s\033[0m\n", finalMsg)
	}
}
