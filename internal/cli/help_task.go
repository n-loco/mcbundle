package cli

import (
	"fmt"
	"slices"
	"strings"

	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/txtui"
)

var helpTask = TaskDefs{
	Name:     "help",
	Aliases:  []string{"--help", "-h"},
	Doc:      "prints this message.",
	Requires: 0,
	Execute: func(*projctx.ProjectContext) {
		txtui.Print(txtui.UIPartOut, "Usage: "+txtui.EscapeUnderline+"bpbuild [task]"+txtui.EscapeReset+"\n\n")
		txtui.Print(txtui.UIPartOut, "Tasks:\n")

		for i, task := range taskList {
			printTaskDoc(task)
			if i < len(taskList)-1 {
				txtui.Print(txtui.UIPartOut, "\n")
			}
		}
	},
}

func printTaskDoc(taskDefs *TaskDefs) {
	name := taskDefs.Name
	aliases := taskDefs.Aliases
	tDoc := taskDefs.Doc

	names := strings.Join(slices.Concat([]string{name}, aliases), ", ")
	names = fmt.Sprintf("  %-20s", names)

	docLines := breakDocInLines(tDoc)

	for i, docLine := range docLines {
		if i == 0 {
			txtui.Printf(txtui.UIPartOut, "%s  %-30s\n", names, docLine)
		} else {
			txtui.Printf(txtui.UIPartOut, "%22s  %-30s\n", "", docLine)
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

		if len(testLine) > 30 {
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
