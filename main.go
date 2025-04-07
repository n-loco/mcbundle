package main

import (
	"github.com/n-loco/bpbuild/internal/cli"
	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/terminal"
)

func main() {
	terminal.SetUpANSIFormatCodes()

	cli.SetupTasks()
	taskDefs := cli.GetTask()

	projCtx := projctx.CreateProjectContext(taskDefs.Requires)

	taskDefs.Execute(&projCtx)
}
