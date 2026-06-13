package editor

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func GetContentFromEditor(initialContent string) (string, error) {
	file, err := os.CreateTemp("", "pasteCTL-*.txt")
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

	editor := os.Getenv("EDITOR")
	if editor == "" {
		if runtime.GOOS == "windows" {
			editor = "notepad"
		} else {
			editor = "vim"
		}
	}

	// Build the argument list. GUI editors like VS Code return immediately
	// unless told to wait, which causes the temp file to be read (empty/unchanged)
	// and deleted before the user can edit it.
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