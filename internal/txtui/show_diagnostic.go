package txtui

import (
	"fmt"

	"github.com/mcbundle/mcbundle/internal/alert"
)

var (
	warnPrefix = EscapeBold + EscapeColorRaw(0xeec520, false) + "[ ! WARNING ]:" + EscapeReset
	errPrefix  = EscapeBold + EscapeColorRaw(0xed3a31, false) + "[ X ERROR ]:" + EscapeReset
	tipPrefix  = EscapeBold + EscapeColorRaw(0x864aba, false) + "[ + TIP ]:" + EscapeReset
)

func ShowDiagnostic(diagnostic alert.Diagnostic) {
	if diagnostic.IsZero() {
		return
	}

	warnCount := len(diagnostic.Warnings)
	errCount := len(diagnostic.Errors)

	if warnCount > 0 {
		showDiagnosticList(warnPrefix, diagnostic.Warnings)
	}
	if errCount > 0 {
		showDiagnosticList(errPrefix, diagnostic.Errors)
	}

	if warnCount > 0 {
		errUIPart.WriteString(fmt.Sprintf("Total warnings: %d\n", warnCount))
	}
	if errCount > 0 {
		errUIPart.WriteString(fmt.Sprintf("Total errors: %d\n", errCount))
	}
}

func showDiagnosticList(prefix string, alerts []alert.Alert) {
	for _, alert := range alerts {
		var display = alert.Display()

		errUIPart.ansiAwareWrite(prefix)
		errUIPart.WriteString(" ")
		errUIPart.ansiAwareWrite(display.Message)

		if tip := display.Tip; len(tip) > 0 {
			errUIPart.WriteString("\n")
			errUIPart.ansiAwareWrite(tipPrefix)
			errUIPart.WriteString(" ")
			errUIPart.ansiAwareWrite(tip)
		}

		errUIPart.WriteString("\n")
	}
}
