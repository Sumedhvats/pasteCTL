package cmd

import (
	"fmt"

	"github.com/Sumedhvats/pastectl/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage CLI configuration",
}

var setCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration key-value pair",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := args[1]
		if err := config.Set(key, value); err != nil {
			fmt.Printf("Error setting config: %v\n", err)
			return
		}
		fmt.Printf(" Config set: %s = %s\n", key, value)
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	rootCmd.AddCommand(configCmd)
}