//go:build unix

package cli

import "os"

func isTerminal(file *os.File) bool {
	return true
}
