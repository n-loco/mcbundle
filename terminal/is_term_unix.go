//go:build unix

package terminal

import (
	"os"

	"golang.org/x/sys/unix"
)

func IsTerminal(file *os.File) bool {
	_, err := unix.IoctlGetWinsize(int(file.Fd()), unix.TIOCGWINSZ)
	return err == nil
}
