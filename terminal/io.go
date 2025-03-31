package terminal

import (
	"fmt"
	"os"
	"strings"
)

type outCfg struct {
	File *os.File
	S1   string
	S2   string
}

var normalCfg = outCfg{
	os.Stdout,
	"",
	"",
}

var errorCfg = outCfg{
	os.Stderr,
	BoldRed + "Error: " + Reset,
	BoldRed + "     : " + Reset,
}

var warnCfg = outCfg{
	os.Stderr,
	BoldYellow + "Warning: " + Reset,
	BoldYellow + "       : " + Reset,
}

var lastOutCfg = &normalCfg

func Printf(format string, a ...any) {
	fprintfInternal(&normalCfg, format, a...)
}

func Print(msg string) {
	fprintInternal(&normalCfg, msg)
}

func Eprintf(format string, a ...any) {
	fprintfInternal(&errorCfg, format, a...)
}

func Eprint(msg string) {
	fprintInternal(&errorCfg, msg)
}

func Wprintf(format string, a ...any) {
	fprintfInternal(&warnCfg, format, a...)
}

func Wprint(msg string) {
	fprintInternal(&warnCfg, msg)
}

func fprintfInternal(selectedOutCfg *outCfg, format string, a ...any) {
	fprintInternal(selectedOutCfg, fmt.Sprintf(format, a...))
}

func fprintInternal(selectedOutCfg *outCfg, rawMsg string) {
	if !IsTerminal(selectedOutCfg.File) {
		rawMsg = StripANSI(rawMsg)
	}

	lines := strings.Split(rawMsg, "\n")

	for i, line := range lines {
		isNotEndLine := i < len(lines)-1

		if isNotEndLine {
			line += "\n"
			if lastOutCfg != selectedOutCfg {
				line = selectedOutCfg.S1 + line
			} else {
				line = selectedOutCfg.S2 + line
			}
		}

		fmt.Fprint(selectedOutCfg.File, line)

		lastOutCfg = selectedOutCfg
	}
}
