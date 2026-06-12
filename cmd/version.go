package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is bumped for each tagged release.
var Version = "1.1.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of pasteCTL",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pastectl v%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
