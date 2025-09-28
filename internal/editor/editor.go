package editor

import (
	"os"
	"os/exec"
	"runtime"
)

func GetContentFromEditor(initialContent string) (string, error) {
	file, err := os.CreateTemp("", "pastectl-*.txt")
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

	cmd := exec.Command(editor, file.Name())
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