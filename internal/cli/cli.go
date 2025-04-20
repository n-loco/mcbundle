package cli

import (
	"os"

	"github.com/n-loco/bpbuild/internal/txtui"
)

var cmdMap = map[string]*commandDefinitions{}
var cmdList = []*commandDefinitions{}

func registerCommand(cmdDefs *commandDefinitions) {
	cmdList = append(cmdList, cmdDefs)
	cmdMap[cmdDefs.name] = cmdDefs
	for _, alias := range cmdDefs.aliases {
		cmdMap[alias] = cmdDefs
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

func getCommand() *commandDefinitions {
	if len(os.Args) < 2 {
		return &helpCmd
	}

	cmdName := os.Args[1]

	cmdDefs, exists := cmdMap[cmdName]

	if !exists {
		txtui.PrePrintf(txtui.UIPartErr, txtui.ErrPrefix, "unknown command: %s\n", cmdName)
		txtui.Printf(txtui.UIPartErr, "use "+txtui.EscapeItalic+"bpbuild help"+txtui.EscapeReset+" to see a list of commands\n")
		os.Exit(1)
	}

	return cmdDefs
}

func Entry() {
	setupCommands()

	cmdDefs := getCommand()

	var optSlice []string

	if len(os.Args) > 2 {
		optSlice = os.Args[2:]
	}

	cmdDefs.execute(optSlice)
}
