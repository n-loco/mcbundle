package txtui

import "fmt"

const (
	EscapeReset = "\x1b[0m"

	EscapeBold      = "\x1b[1m"
	EscapeItalic    = "\x1b[3m"
	EscapeUnderline = "\x1b[4m"
)

func EscapeDefaultColor(background bool) string {
	if background {
		return "\x1b[49m"
	}
	return "\x1b[39m"
}

func EscapeColorRaw(color uint32, background bool) string {
	r := uint8((color & 0xff0000) >> 16)
	g := uint8((color & 0x00ff00) >> 8)
	b := uint8(color & 0x0000ff)

	return EscapeColorRGB(r, g, b, background)
}

func EscapeColorRGB(r, g, b uint8, background bool) string {
	var tn string
	if background {
		tn = "48"
	} else {
		tn = "38"
	}

	return fmt.Sprintf("\x1b[%s;2;%d;%d;%dm", tn, r, g, b)
}
