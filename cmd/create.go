package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Sumedhvats/pastectl/internal/api"
	"github.com/Sumedhvats/pastectl/internal/config"
	"github.com/Sumedhvats/pastectl/internal/editor"
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
	createCmd.Flags().StringVarP(&expire, "expire", "e", "never", "Set expiration time (e.g., 10m, 1h)")
	rootCmd.AddCommand(createCmd)
}

func createPaste(cmd *cobra.Command, args []string) {
	var content string
	var err error
fmt.Print(args)
	// 1. Determine Content Source (File vs. Editor)
	if filePath != "" {
		fileContent, fileErr := os.ReadFile(filePath)
		if fileErr != nil {
			log.Fatalf("❌ Error reading file: %v", fileErr)
		}
		content = string(fileContent)
	} else {
		content, err = editor.GetContentFromEditor("")
		if err != nil {
			log.Fatalf("❌ Error opening editor: %v", err)
		}
	}

	// 2. Handle Empty Content
	if strings.TrimSpace(content) == "" {
		fmt.Println("⚠️ No content provided. Aborting creation.")
		return
	}

	// 3. Determine Language
	if language == "" && filePath != "" {
		detectedLang := mapExtensionToLanguage(filepath.Ext(filePath))
		fmt.Printf("ℹ️ Detected language: %s\n", detectedLang)
		language = detectedLang
	} else if language == "" {
		language = "plain" // Default for editor input
	}

	// 4. Call API and Handle Errors
	fmt.Println("🚀 Creating paste...")
	paste, err := api.CreatePaste(content, language, expire)
	if err != nil {
		log.Fatalf("❌ Failed to create paste: %v", err)
	}

	// 5. Provide Clear Success Output
	frontendURL := config.Get("frontend_url")
	if frontendURL == "" {
		log.Fatalf("Error: frontend_url is not set. Please use 'pastectl config set frontend_url <url>'")
	}

	fmt.Printf("✅ Paste created successfully!\n")
	fmt.Printf("🔗 Link: %s/%s\n", frontendURL, paste.ID)
}

func mapExtensionToLanguage(ext string) string {
	trimmedExt := strings.TrimPrefix(strings.ToLower(ext), ".")
	langMap := map[string]string{
		"js":   "javascript",
		"py":   "python",
		"go":   "go",
		"java": "java",
		"c":    "c",
		"cpp":  "cpp",
		"json": "json",
		"md":   "markdown",
		"txt":  "plain",
	}
	if lang, ok := langMap[trimmedExt]; ok {
		return lang
	}
	return "plain"
}