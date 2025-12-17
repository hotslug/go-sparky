package logger

import (
	_ "embed"
	"fmt"
)

//go:embed sparky.txt
var sparkyBanner string

// PrintBanner renders the ASCII banner to stdout.
func PrintBanner() {
	if sparkyBanner == "" {
		fmt.Println("go-sparky")
		return
	}

	fmt.Println(sparkyBanner)
}
