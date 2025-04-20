package cli

import (
	"os"

	"github.com/n-loco/bpbuild/internal/txtui"
)

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

func getCommand() command {
	if len(os.Args) < 2 {
		return &helpCmd
	}

	cmdName := os.Args[1]

	cmd, exists := cmdMap[cmdName]

	if !exists {
		txtui.PrePrintf(txtui.UIPartErr, txtui.ErrPrefix, "unknown command: %s\n", cmdName)
		txtui.Printf(txtui.UIPartErr, "use "+txtui.EscapeItalic+"bpbuild help"+txtui.EscapeReset+" to see a list of commands\n")
		os.Exit(1)
	}

	return cmd
}

func Entry() {
	setupCommands()

	cmd := getCommand()

	var optSlice []string

	if len(os.Args) > 2 {
		optSlice = os.Args[2:]
	}

	cmd.execute(optSlice)
}
