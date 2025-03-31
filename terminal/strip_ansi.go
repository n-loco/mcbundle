package terminal

import (
	"regexp"
)

var stripANSIRegExp = regexp.MustCompile(`(\x{1b}\[.*?m)`)

func StripANSI(str string) string {
	return stripANSIRegExp.ReplaceAllLiteralString(str, "")
}
