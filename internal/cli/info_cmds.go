package cli

import (
	"fmt"
	"slices"
	"strings"

	"github.com/n-loco/bpbuild/internal/assets"
	"github.com/n-loco/bpbuild/internal/txtui"
)

type versionCommand empty

var versionCmd = versionCommand{}

var versionCmdInfo = commandInfo{
	name:    "version",
	aliases: []string{"--version", "-v"},
	doc:     "...",
}

func (cmd versionCommand) info() *commandInfo {
	return &versionCmdInfo
}

func (cmd versionCommand) execute([]string) {
	txtui.Printf(txtui.UIPartOut, "%s\n", assets.ProgramVersion)
}

type helpCommand empty

var helpCmd = helpCommand{}

var helpCmdInfo = commandInfo{
	name:    "help",
	aliases: []string{"--help", "-h"},
	doc:     "...",
}

func (cmd helpCommand) info() *commandInfo {
	return &helpCmdInfo
}

func (cmd helpCommand) execute([]string) {
	txtui.Print(txtui.UIPartOut, "Usage: "+txtui.EscapeItalic+"bpbuild [command] <options>"+txtui.EscapeReset+"\n\n")
	txtui.Print(txtui.UIPartOut, "Commands:\n")

	for i, cmd := range cmdList {
		printCommandDoc(cmd)
		if i < len(cmdList)-1 {
			txtui.Print(txtui.UIPartOut, "\n")
		}
	}
}

func printCommandDoc(cmdDefs command) {
	cmdInfo := cmdDefs.info()
	name := cmdInfo.name
	aliases := cmdInfo.aliases
	tDoc := cmdInfo.doc

	names := strings.Join(slices.Concat([]string{name}, aliases), ", ")
	names = fmt.Sprintf("  %-28s", names)

	docLines := breakDocInLines(tDoc)

	for i, docLine := range docLines {
		if i == 0 {
			txtui.Printf(txtui.UIPartOut, "%s  %-35s\n", names, docLine)
		} else {
			txtui.Printf(txtui.UIPartOut, "%30s  %-35s\n", "", docLine)
		}
	}
}

func breakDocInLines(docStr string) []string {
	var lines []string

	words := strings.Split(docStr, " ")

	newLine := ""
	for _, word := range words {
		var testLine string

		if len(newLine) == 0 {
			testLine = word
		} else {
			testLine = strings.Join([]string{newLine, word}, " ")
		}

		if len(testLine) > 35 {
			lines = append(lines, newLine)
			newLine = word
		} else {
			newLine = testLine
		}
	}

	if len(newLine) > 0 {
		lines = append(lines, newLine)
	}

	return lines
}
