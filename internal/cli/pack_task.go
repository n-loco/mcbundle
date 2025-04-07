package cli

import (
	"github.com/n-loco/bpbuild/internal/operations"
	"github.com/n-loco/bpbuild/internal/projctx"
)

var packTask = TaskDefs{
	Requires: projctx.EnvRequireFlagRecipe,
	Name:     "pack",
	Doc:      "...",
	Execute: func(projctx *projctx.ProjectContext) {
		operations.PackProject(projctx)
	},
}
