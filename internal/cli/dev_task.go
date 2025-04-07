package cli

import (
	"github.com/n-loco/bpbuild/internal/operations"
	"github.com/n-loco/bpbuild/internal/projctx"
)

var devTask = TaskDefs{
	Requires: projctx.EnvRequireFlagRecipe | projctx.EnvRequireFlagComMojang,
	Name:     "dev",
	Doc:      "...",
	Execute: func(projctx *projctx.ProjectContext) {
		operations.CopyToDev(projctx)
	},
}
