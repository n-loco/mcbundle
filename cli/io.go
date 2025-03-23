package cli

import (
	"fmt"
	"os"

	"github.com/redrock/autocrafter/ansi"
)

const (
	errorPrefix   = ansi.BoldRed + "Error: " + ansi.Reset
	warningPrefix = ansi.BoldYellow + "Warning: " + ansi.Reset
)

func Printf(format string, a ...any) {
	fprintfInternal(os.Stdout, format, a...)
}

func Print(msg string) {
	fprintInternal(os.Stdout, msg)
}

func Eprintf(format string, a ...any) {
	fprintfInternal(os.Stderr, errorPrefix+format, a...)
}

func Eprint(msg string) {
	fprintInternal(os.Stderr, errorPrefix+msg)
}

func Wprintf(format string, a ...any) {
	fprintfInternal(os.Stderr, warningPrefix+format, a...)
}

func Wprint(msg string) {
	fprintInternal(os.Stderr, warningPrefix+msg)
}

func fprintfInternal(file *os.File, format string, a ...any) {
	fprintInternal(file, fmt.Sprintf(format, a...))
}

func fprintInternal(file *os.File, msg string) {
	fmt.Fprint(file, ansi.StripANSIWhenFile(file, msg))
}
