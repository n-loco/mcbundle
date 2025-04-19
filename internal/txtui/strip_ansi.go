package txtui

import (
	"regexp"
)

var stripANSIRegExp = regexp.MustCompile(`(\x{1b}\[.*?m)`)

func stripANSI(str string) string {
	return stripANSIRegExp.ReplaceAllLiteralString(str, "")
}
