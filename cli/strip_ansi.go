package cli

import (
	"os"
	"regexp"
)

var stripANSIRegExp = regexp.MustCompile(`(\x{1b}\[.*?m)`)

func stripANSIIfIsNotTerminal(file *os.File, str string) string {
	if isTerminal(file) {
		return str
	}

	return stripANSIRegExp.ReplaceAllLiteralString(str, "")
}
