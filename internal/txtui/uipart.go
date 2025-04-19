package txtui

import "os"

type UIOPart uint8

const (
	UIPartOut UIOPart = 1 + iota
	UIPartErr
)

func (pType UIOPart) part() uiPart {
	switch pType {
	case UIPartOut:
		return outUIPart
	case UIPartErr:
		return errUIPart
	default:
		panic("invalid txtui part.")
	}
}

type uiPart struct {
	*os.File
	terminal bool
}

var outUIPart = createUIPart(os.Stdout)
var errUIPart = createUIPart(os.Stderr)

func (part uiPart) ansiAwareWrite(s string) (n int, err error) {
	if !part.terminal {
		s = stripANSI(s)
	}
	n, err = part.WriteString(s)
	return
}

func createUIPart(osFile *os.File) uiPart {
	var part uiPart
	part.File = osFile
	part.terminal = isTerminal(osFile)
	setupANSICodes(part)
	return part
}
