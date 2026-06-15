package editor

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// languageExtensions maps language names to file extensions.
var languageExtensions = map[string]string{
	"javascript":  ".js",
	"typescript":  ".ts",
	"python":      ".py",
	"go":          ".go",
	"rust":        ".rs",
	"java":        ".java",
	"c":           ".c",
	"cpp":         ".cpp",
	"csharp":      ".cs",
	"ruby":        ".rb",
	"php":         ".php",
	"swift":       ".swift",
	"kotlin":      ".kt",
	"scala":       ".scala",
	"html":        ".html",
	"css":         ".css",
	"scss":        ".scss",
	"less":        ".less",
	"json":        ".json",
	"yaml":        ".yaml",
	"yml":         ".yml",
	"toml":        ".toml",
	"xml":         ".xml",
	"markdown":    ".md",
	"sql":         ".sql",
	"shell":       ".sh",
	"bash":        ".sh",
	"powershell":  ".ps1",
	"dockerfile":  ".dockerfile",
	"lua":         ".lua",
	"perl":        ".pl",
	"r":           ".r",
	"dart":        ".dart",
	"elixir":      ".ex",
	"erlang":      ".erl",
	"haskell":     ".hs",
	"clojure":     ".clj",
	"jsx":         ".jsx",
	"tsx":         ".tsx",
	"vue":         ".vue",
	"svelte":      ".svelte",
	"graphql":     ".graphql",
	"protobuf":    ".proto",
	"plaintext":   ".txt",
	"text":        ".txt",
}

// extForLanguage returns the file extension for a language name.
func extForLanguage(language string) string {
	if language == "" {
		return ".txt"
	}
	if ext, ok := languageExtensions[strings.ToLower(language)]; ok {
		return ext
	}
	// If the language looks like it could already be an extension, use it directly
	if !strings.Contains(language, " ") && len(language) <= 10 {
		return "." + strings.ToLower(language)
	}
	return ".txt"
}

func GetContentFromEditor(initialContent string, editorOverride string, language string) (string, error) {
	ext := extForLanguage(language)
	file, err := os.CreateTemp("", "pasteCTL-*"+ext)
	if err != nil {
		return "", err
	}
	defer os.Remove(file.Name())

	if len(initialContent) > 0 {
		if _, err := file.Write([]byte(initialContent)); err != nil {
			return "", err
		}
	}

	if err := file.Close(); err != nil {
		return "", err
	}

	editor := editorOverride
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}
	if editor == "" {
		if runtime.GOOS == "windows" {
			editor = "notepad"
		} else {
			editor = "vim"
		}
	}

	args := []string{}
	editorBase := filepath.Base(editor)
	switch editorBase {
	case "code", "code-insiders":
		args = append(args, "--wait")
	case "subl", "sublime_text":
		args = append(args, "--wait")
	case "gedit":
		args = append(args, "--wait")
	}
	args = append(args, file.Name())

	cmd := exec.Command(editor, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	contentBytes, err := os.ReadFile(file.Name())
	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}