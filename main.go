package main

import (
	"github.com/n-loco/bpbuild/internal/cli"
	"github.com/n-loco/bpbuild/internal/projctx"
)

func main() {
	cli.SetupTasks()
	taskDefs := cli.GetTask()

	projCtx := projctx.CreateProjectContext(taskDefs.Requires)

	taskDefs.Execute(&projCtx)
}
