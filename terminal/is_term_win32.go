//go:build windows

package terminal

import (
	"os"

	"golang.org/x/sys/windows"
)

func IsTerminal(file *os.File) bool {
	var _m uint32
	err := windows.GetConsoleMode(windows.Handle(file.Fd()), &_m)
	return err == nil
}
