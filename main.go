package main

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/cli"
	"github.com/mcbundle/mcbundle/internal/txtui"
)

func main() {
	var diagnostic = alert.NewDiagnostic()
	cli.Entry(diagnostic)

	txtui.ShowDiagnostic(diagnostic)

	if diagnostic.HasErrors() {
		os.Exit(-1)
	}
}
