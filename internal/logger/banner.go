package logger

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrintBanner renders the ASCII banner to stdout.
func PrintBanner() {
	path := filepath.Join("assets", "ascii", "sparky.txt")
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("go-sparky")
		return
	}

	fmt.Println(string(data))
}
