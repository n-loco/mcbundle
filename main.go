package main

import (
	"os"

	"github.com/n-loco/bpbuild/internal/cli"
	"github.com/n-loco/bpbuild/internal/txtui"
)

func main() {
	diagnostic := cli.Entry()

	txtui.ShowDiagnostic(diagnostic)

	if diagnostic.HasErrors() {
		os.Exit(-1)
	}
}
