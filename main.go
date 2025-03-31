package main

import (
	"github.com/redrock/autocrafter/cli"
	"github.com/redrock/autocrafter/terminal"
)

func main() {
	terminal.SetUpANSIFormatCodes()

	cli.SetupTasks()
	taskDefs := cli.GetTask()

	dependencies := cli.GetEnvironmentDependencies(taskDefs.Dependencies)
	taskDefs.Execute(&dependencies)
}
