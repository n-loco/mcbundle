//go:build windows

package txtui

import (
	"golang.org/x/sys/windows"
)

func setupANSICodes(part uiPart) {
	if !part.terminal {
		return
	}

	var handle windows.Handle = windows.Handle(part.Fd())
	var mode uint32
	windows.GetConsoleMode(handle, &mode)

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	windows.SetConsoleMode(handle, mode)
}
