package cmd

import "github.com/spf13/cobra"

const appVersion = "0.2.0"

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the go-sparky version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(appVersion)
		},
	}
}
