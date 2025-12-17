package cmd

import (
	"fmt"
	"os"

	"github.com/hotslug/go-sparky/internal/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-sparky",
	Short: "Scaffold a modern React app with Vite",
}

var flagVerbose bool

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		logger.SetVerbose(flagVerbose)
	}

	rootCmd.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", true, "Enable verbose output (spinners, extra logs)")
	rootCmd.AddCommand(newNewCmd())
	rootCmd.AddCommand(newVersionCmd())
}
