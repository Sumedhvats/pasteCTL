package utils

import "regexp"

const ansi = "[\x1b\x9b][\\[()#;?]*(?:[0-9]{1,4}(?:;[0-9]{0,4})*)?[0-9A-ORZcf-nqry=><]"

var re = regexp.MustCompile(ansi)

// StripANSI removes ANSI escape codes from a string.
func StripANSI(str string) string {
	return re.ReplaceAllString(str, "")
}
