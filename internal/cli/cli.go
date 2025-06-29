package cli

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/txtui"
)

type UnknownCommandErrorAlert struct {
	CommandName string
}

func (errAlert UnknownCommandErrorAlert) Display() alert.AlertDisplay {
	return alert.AlertDisplay{
		Message: "unknown command: " + errAlert.CommandName,
		Tip:     "use " + txtui.EscapeItalic + "mcbundle help" + txtui.EscapeReset + " to see a list of commands",
	}
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

func getCommand(diagnostic alert.Diagnostic) (cmd command) {
	if len(os.Args) < 2 {
		cmd = &helpCmd
		return
	}

	var cmdName = os.Args[1]
	var cmdExists bool

	cmd, cmdExists = cmdMap[cmdName]

	if !cmdExists {
		diagnostic.AppendError(&UnknownCommandErrorAlert{CommandName: cmdName})
	}

	return
}

func Entry(diagnostic alert.Diagnostic) {
	setupCommands()

	var cmd = getCommand(diagnostic)

	if diagnostic.HasErrors() {
		return
	}

	var optSlice []string

	if len(os.Args) > 2 {
		optSlice = os.Args[2:]
	}

	cmd.execute(optSlice, diagnostic)
}
