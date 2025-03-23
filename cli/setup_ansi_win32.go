//go:build windows

package cli

import (
	"os"

	"golang.org/x/sys/windows"
)

func SetUpANSIEscapeCodes() {
	enableVirtualTerminalProcessing(windows.Handle(os.Stdout.Fd()))
	enableVirtualTerminalProcessing(windows.Handle(os.Stderr.Fd()))
}

func enableVirtualTerminalProcessing(handle windows.Handle) {
	var mode uint32
	err := windows.GetConsoleMode(handle, &mode)

	if err == nil {
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		windows.SetConsoleMode(handle, mode)
	}
}
