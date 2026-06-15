package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"github.com/mdp/qrterminal/v3"
	"github.com/Sumedhvats/pasteCTL/internal/api"
	"github.com/Sumedhvats/pasteCTL/internal/config"
	"github.com/Sumedhvats/pasteCTL/internal/editor"
	"github.com/Sumedhvats/pasteCTL/internal/utils"
	"github.com/spf13/cobra"
)

var (
	filePath   string
	language   string
	expire     string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new paste from a file or editor",
	Long: `Creates a new paste.

You can provide content in one of two ways:
1. By specifying a file path with the --file flag.
2. By launching a text editor (default behavior).`,
	Run: createPaste,
}

func init() {
	createCmd.Flags().StringVarP(&filePath, "file", "f", "", "Create paste from a file path")
	createCmd.Flags().StringVarP(&language, "language", "l", "", "Override language detection (e.g., go, python)")
	createCmd.Flags().StringVarP(&expire, "expire", "e", "1h", "Set expiration time (e.g., 1h, 24h, 7d, never)")
	rootCmd.AddCommand(createCmd)
}

func createPaste(cmd *cobra.Command, args []string) {
	var content string
	var err error


	if filePath != "" {
		fileContent, fileErr := os.ReadFile(filePath)
		if fileErr != nil {
			log.Fatalf("Error reading file: %v", fileErr)
		}
		content = string(fileContent)
	} else {
		content, err = editor.GetContentFromEditor("", "", language)
		if err != nil {
			log.Fatalf("Error opening editor: %v", err)
		}
	}
	if strings.TrimSpace(content) == "" {
		fmt.Println("No content provided. Aborting creation.")
		return
	}

	content = utils.StripANSI(content)

	if language == "" && filePath != "" {
		detectedLang := mapExtensionToLanguage(filepath.Ext(filePath))
		fmt.Printf("Detected language: %s\n", detectedLang)
		language = detectedLang
	} else if language == "" {
		language = "plain"
	}
	fmt.Println("Creating paste...")
	paste, err := api.CreatePaste(content, language, expire)
	if err != nil {
		log.Fatalf("Failed to create paste: %v", err)
	}
	frontendURL := config.Get("frontend_url")
	if frontendURL == "" {
		log.Fatalf("Error: frontend_url is not set. Please use 'pasteCTL config set frontend_url <url>'")
	}
	config := qrterminal.Config{
		HalfBlocks: true,          
		Level:      qrterminal.L,   
		Writer:     os.Stdout,
		QuietZone:  1,             
	}
	fmt.Printf("Paste created successfully!\n")
	fmt.Printf("Link: %s/%s\n", frontendURL, paste.ID)
	qrUrl:= fmt.Sprintf("%s/%s", frontendURL, paste.ID)
	qrterminal.GenerateWithConfig(qrUrl, config)
	fmt.Println()

}

func mapExtensionToLanguage(ext string) string {
	trimmedExt := strings.TrimPrefix(strings.ToLower(ext), ".")
	langMap := map[string]string{

		"js":  "javascript",
		"jsx": "javascript",
		"ts":  "javascript",
		"tsx": "javascript",
		"mjs": "javascript",
		"cjs": "javascript",

		"py":  "python",
		"pyw": "python",

		"java": "java",
		"cpp": "cpp",
		"cc":  "cpp",
		"cxx": "cpp",
		"hpp": "cpp",
		"hxx": "cpp",
		"h":   "cpp",

		"c": "c",

		"go": "go",

		"sql": "sql",
		// Plain text fallbacks
		"txt":  "plain",
		"md":   "plain",
		"json": "plain",
	}
	if lang, ok := langMap[trimmedExt]; ok {
		return lang
	}
	return "plain"
}