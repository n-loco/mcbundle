package main

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/cli"
	"github.com/mcbundle/mcbundle/internal/txtui"
)

func main() {
	diagnostic := cli.Entry()

	txtui.ShowDiagnostic(diagnostic)

	if diagnostic.HasErrors() {
		os.Exit(-1)
	}
}
