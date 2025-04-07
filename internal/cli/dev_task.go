package cli

import (
	"github.com/n-loco/mcbuild/internal/operations"
	"github.com/n-loco/mcbuild/internal/projctx"
)

var devTask = TaskDefs{
	Requires: projctx.EnvRequireFlagRecipe | projctx.EnvRequireFlagComMojang,
	Name:     "dev",
	Doc:      "...",
	Execute: func(projctx *projctx.ProjectContext) {
		operations.CopyToDev(projctx)
	},
}
