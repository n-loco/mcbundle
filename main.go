package main

import (
	"github.com/n-loco/mcbuild/internal/cli"
	"github.com/n-loco/mcbuild/internal/terminal"
)

func main() {
	terminal.SetUpANSIFormatCodes()

	cli.SetupTasks()
	taskDefs := cli.GetTask()

	dependencies := cli.GetEnvironmentDependencies(taskDefs.Dependencies)
	taskDefs.Execute(&dependencies)
}
