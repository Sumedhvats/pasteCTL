package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is bumped for each tagged release.
var Version = "1.1.4"
var Commit = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of pasteCTL",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pasteCTL v%s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
