package txtui

import "fmt"

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
