/*
Copyright © 2025 SUMEDH VATS sumedhvats2004@gmail.com
*/
package cmd

import (
	"os"

	"github.com/Sumedhvats/pastectl/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pastectl",
	Short: "Share code directly on the pasteCTL platform",
	Long: `pasteCTL is a modern, full-stack pastebin service for storing, sharing, and managing code snippets and text. It's built as a monorepo with a robust Go backend, a sleek Next.js frontend, and a command-line interface (CLI).`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


