package cmd

import (
	"fmt"
	"log"

	"github.com/Sumedhvats/pastectl/internal/api"
	"github.com/spf13/cobra"
)

var raw bool

var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a paste by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pasteID := args[0]

		if raw {
			content, err := api.GetPasteRaw(pasteID)
			if err != nil {
				log.Fatalf("❌ Failed to get raw paste: %v", err)
			}
			fmt.Println(content)
		} else {
			paste, err := api.GetPaste(pasteID)
			if err != nil {
				log.Fatalf("❌ Failed to get paste: %v", err)
			}
			fmt.Printf("--- Paste Details ---\n")
			fmt.Printf("ID:       %s\n", paste.ID)
			fmt.Printf("Language: %s\n", paste.Language)
			fmt.Printf("Created:  %s\n", paste.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Printf("--- Content ---\n%s\n", paste.Content)
		}
	},
}

func init() {
	getCmd.Flags().BoolVar(&raw, "raw", false, "Display only the raw content of the paste")
	rootCmd.AddCommand(getCmd)
}