package ansi

import (
	"os"
	"regexp"

	"github.com/redrock/autocrafter/terminal"
)

var stripANSIRegExp = regexp.MustCompile(`(\x{1b}\[.*?m)`)

func StripANSIWhenFile(file *os.File, str string) string {
	if terminal.IsTerminal(file) {
		return str
	}

	return stripANSIRegExp.ReplaceAllLiteralString(str, "")
}
