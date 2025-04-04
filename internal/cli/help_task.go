package cli

import (
	"fmt"
	"slices"
	"strings"

	"github.com/n-loco/mcbuild/internal/terminal"
)

var helpTask = TaskDefs{
	Name:         "help",
	Aliases:      []string{"--help", "-h"},
	Doc:          "prints this message.",
	Dependencies: 0,
	Execute: func(*EnvironmentDependencies) {
		terminal.Print("Usage: " + terminal.UnderlineWhite + "autocrafter [task]" + terminal.Reset + "\n\n")
		terminal.Print("Tasks:\n")

		for i, task := range taskList {
			printTaskDoc(task)
			if i < len(taskList)-1 {
				terminal.Print("\n")
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
			terminal.Printf("%s  %-30s\n", names, docLine)
		} else {
			terminal.Printf("%22s  %-30s\n", "", docLine)
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
