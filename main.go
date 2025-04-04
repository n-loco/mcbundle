package main

import (
	"github.com/redrock/autocrafter/internal/cli"
	"github.com/redrock/autocrafter/internal/terminal"
)

func main() {
	terminal.SetUpANSIFormatCodes()

	cli.SetupTasks()
	taskDefs := cli.GetTask()

	dependencies := cli.GetEnvironmentDependencies(taskDefs.Dependencies)
	taskDefs.Execute(&dependencies)
}
