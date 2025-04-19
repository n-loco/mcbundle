package txtui

import "fmt"

var (
	ErrPrefix  = EscapeBold + EscapeColorRaw(0xed3a31, false) + "[ X ERROR ]:" + EscapeReset
	WarnPrefix = EscapeBold + EscapeColorRaw(0xeec520, false) + "[ ! WARNING ]:" + EscapeReset
)

func Print(pType UIOPart, s string) (n int, err error) {
	return pType.part().ansiAwareWrite(s)
}

func PrePrint(pType UIOPart, pre, s string) (n int, err error) {
	return pType.part().ansiAwareWrite(pre + " " + s)
}

func Printf(pType UIOPart, format string, a ...any) (n int, err error) {
	return pType.part().ansiAwareWrite(fmt.Sprintf(format, a...))
}

func PrePrintf(pType UIOPart, pre, format string, a ...any) (n int, err error) {
	return pType.part().ansiAwareWrite(pre + " " + fmt.Sprintf(format, a...))
}
