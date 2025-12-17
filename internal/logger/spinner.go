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
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Printf("\r%s %s", frames[i%len(frames)], msg)
				i++
				time.Sleep(120 * time.Millisecond)
			}
		}
	}()

	return func(finalMsg string) {
		close(stop)
		fmt.Printf("\r✓ %s\n", finalMsg)
	}
}
