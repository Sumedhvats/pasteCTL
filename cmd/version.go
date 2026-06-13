package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "1.1.6"
var Commit = "dev"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of pasteCTL",
	Run: func(cmd *cobra.Command, args []string) {
		if Commit != "dev" {
			fmt.Printf("pasteCTL v%s (%s)\n", Version, Commit)
		} else {
			fmt.Printf("pasteCTL v%s\n", Version)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
