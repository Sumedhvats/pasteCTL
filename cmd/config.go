package cmd

import (
	"fmt"

	"github.com/Sumedhvats/pasteCTL/internal/config"
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
		fmt.Printf("Config set: %s = %s\n", key, value)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configuration values",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Config file: %s\n\n", config.ConfigFilePath())
		fmt.Printf("  %-20s %-40s %s\n", "KEY", "VALUE", "SOURCE")
		fmt.Printf("  %-20s %-40s %s\n", "---", "-----", "------")
		values := config.List()
		for _, key := range config.KnownKeys {
			source := "default"
			if !config.IsDefault(key) {
				source = "user-set"
			}
			fmt.Printf("  %-20s %-40s %s\n", key, values[key], source)
		}
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(listCmd)
	rootCmd.AddCommand(configCmd)
}