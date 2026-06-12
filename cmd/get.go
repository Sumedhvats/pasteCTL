package cmd

import (
	"fmt"
	"log"

	"github.com/Sumedhvats/pasteCTL/internal/api"
	"github.com/Sumedhvats/pasteCTL/internal/highlight"
	"github.com/spf13/cobra"
)

var (
	raw     bool
	noColor bool
)

var getCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a paste by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pasteID := args[0]

		if raw {
			content, err := api.GetPasteRaw(pasteID)
			if err != nil {
				log.Fatalf("Failed to get raw paste: %v", err)
			}
			fmt.Println(content)
		} else {
			paste, err := api.GetPaste(pasteID)
			if err != nil {
				log.Fatalf("Failed to get paste: %v", err)
			}
			fmt.Printf("--- Paste Details ---\n")
			fmt.Printf("ID:       %s\n", paste.ID)
			fmt.Printf("Language: %s\n", paste.Language)
			fmt.Printf("Created:  %s\n", paste.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Printf("--- Content ---\n")

			if noColor {
				fmt.Println(paste.Content)
			} else {
				if err := highlight.Print(paste.Content, paste.Language); err != nil {
					// Fall back to plain output if highlighting fails
					fmt.Println(paste.Content)
				}
				fmt.Println() // newline after highlighted output
			}
		}
	},
}

func init() {
	getCmd.Flags().BoolVar(&raw, "raw", false, "Display only the raw content of the paste")
	getCmd.Flags().BoolVar(&noColor, "no-color", false, "Disable syntax highlighting")
	rootCmd.AddCommand(getCmd)
}