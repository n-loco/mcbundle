package cli

import (
	"fmt"
	"os"
)

func Printf(format string, a ...any) {
	fprintfInternal(os.Stdout, format, a...)
}

func Print(msg string) {
	fprintInternal(os.Stdout, msg)
}

func Eprintf(format string, a ...any) {
	fprintfInternal(os.Stderr, format, a...)
}

func Eprint(msg string) {
	fprintInternal(os.Stderr, msg)
}

func Wprintf(format string, a ...any) {
	fprintfInternal(os.Stderr, format, a...)
}

func Wprint(msg string) {
	fprintInternal(os.Stderr, msg)
}

func fprintfInternal(file *os.File, format string, a ...any) {
	fprintInternal(file, fmt.Sprintf(format, a...))
}

func fprintInternal(file *os.File, msg string) {
	fmt.Fprint(file, stripANSIIfIsNotTerminal(file, msg))
}
