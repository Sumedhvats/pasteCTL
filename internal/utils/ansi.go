package utils

import "regexp"

const ansi = `\x1b\[[0-9;]*[a-zA-Z]`

var re = regexp.MustCompile(ansi)

// StripANSI removes ANSI escape codes from a string.
func StripANSI(str string) string {
	return re.ReplaceAllString(str, "")
}
