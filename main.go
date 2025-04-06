package main

import (
	"github.com/n-loco/mcbuild/internal/cli"
	"github.com/n-loco/mcbuild/internal/projctx"
	"github.com/n-loco/mcbuild/internal/terminal"
)

func main() {
	terminal.SetUpANSIFormatCodes()

	cli.SetupTasks()
	taskDefs := cli.GetTask()

	projCtx := projctx.CreateProjectContext(taskDefs.Requires)

	taskDefs.Execute(&projCtx)
}
