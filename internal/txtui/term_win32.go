//go:build windows

package txtui

import (
	"os"

	"golang.org/x/sys/windows"
)

func isTerminal(file *os.File) bool {
	var _m uint32
	err := windows.GetConsoleMode(windows.Handle(file.Fd()), &_m)
	return err == nil
}
