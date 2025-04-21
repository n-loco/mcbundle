package cli

import (
	"os"

	"github.com/n-loco/bpbuild/internal/alert"
	"github.com/n-loco/bpbuild/internal/txtui"
)

type UnknownCommandErrorAlert struct {
	CommandName string
}

func (errAlert UnknownCommandErrorAlert) Display() string {
	return "unknown command: " + errAlert.CommandName
}

func (errAlert UnknownCommandErrorAlert) Tip() string {
	return "use " + txtui.EscapeItalic + "bpbuild help" + txtui.EscapeReset + " to see a list of commands"
}

var cmdMap = map[string]command{}
var cmdList = []command{}

func registerCommand(cmd command) {
	cmdInfo := cmd.info()
	cmdList = append(cmdList, cmd)
	cmdMap[cmdInfo.name] = cmd
	for _, alias := range cmdInfo.aliases {
		cmdMap[alias] = cmd
	}
}

func setupCommands() {
	registerCommand(&devCmd)
	registerCommand(&packCmd)
	registerCommand(&buildCmd)
	registerCommand(&versionCmd)

	// special case
	registerCommand(&helpCmd)
	cmdMap["-?"] = &helpCmd
	cmdMap["/?"] = &helpCmd
	cmdMap["h"] = &helpCmd
}

func getCommand() (cmd command, diagnostic *alert.Diagnostic) {
	if len(os.Args) < 2 {
		cmd = &helpCmd
		return
	}

	cmdName := os.Args[1]
	cmd, exists := cmdMap[cmdName]

	if !exists {
		diagnostic = diagnostic.AppendError(&UnknownCommandErrorAlert{CommandName: cmdName})
	}

	return
}

func Entry() (diagnostic *alert.Diagnostic) {
	setupCommands()

	cmd, getCmdDiag := getCommand()

	diagnostic = diagnostic.Append(getCmdDiag)
	if diagnostic.HasErrors() {
		return
	}

	var optSlice []string

	if len(os.Args) > 2 {
		optSlice = os.Args[2:]
	}

	diagnostic = diagnostic.Append(cmd.execute(optSlice))

	return
}
