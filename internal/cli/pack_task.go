package cli

import (
	"github.com/n-loco/mcbuild/internal/operations"
	"github.com/n-loco/mcbuild/internal/projctx"
)

var packTask = TaskDefs{
	Requires: projctx.EnvRequireFlagRecipe,
	Name:     "pack",
	Doc:      "...",
	Execute: func(projctx *projctx.ProjectContext) {
		operations.PackProject(projctx)
	},
}
