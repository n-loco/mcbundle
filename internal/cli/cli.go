package cli

import (
	"github.com/mcbundle/mcbundle/internal/alert"
)

type entryCommand = func(*argvIterator, alert.Diagnostic)

var entryCommands = map[string]entryCommand{
	// build command
	"build":  buildCmd,
	"bundle": buildCmd,

	// dev command
	"dev":          devCmd,
	"local-deploy": devCmd,

	// pack command
	"pack": packCmd,
	"dist": packCmd,

	// help command
	"help":   helpCmd,
	"--help": helpCmd,
	"-h":     helpCmd,
	"/?":     helpCmd,

	// version command
	"version":   versionCmd,
	"--version": versionCmd,
	"-v":        versionCmd,
}

func Entry(diagnostic alert.Diagnostic) {
	var argv = newArgvIterator()

	if argv.hasNext() {
		var arg = argv.consume()

		var funcCmd, has = entryCommands[arg]
		if has {
			funcCmd(argv, diagnostic)
		} else {
			diagnostic.AppendError(alert.AlertF("unknown command: %s", arg))
		}
	} else {
		helpCmd(argv, diagnostic)
	}
}
