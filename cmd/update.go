package cmd

import (
	"fmt"
	"log"

	"github.com/Sumedhvats/pastectl/internal/api"
	"github.com/Sumedhvats/pastectl/internal/editor"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a paste by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pasteID := args[0]

		existingPaste, err := api.GetPaste(pasteID)
		if err != nil {
			log.Fatalf("❌ Failed to retrieve existing paste: %v", err)
		}

		newContent, err := editor.GetContentFromEditor(existingPaste.Content)
		if err != nil {
			log.Fatalf("❌ Could not get content from editor: %v", err)
		}

		_, err = api.UpdatePaste(pasteID, newContent, existingPaste.Language)
		if err != nil {
			log.Fatalf("❌ Failed to update paste: %v", err)
		}

		fmt.Println("✅ Paste updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}