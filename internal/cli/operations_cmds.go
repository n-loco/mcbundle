package cli

import (
	"github.com/n-loco/bpbuild/internal/operations"
	"github.com/n-loco/bpbuild/internal/projctx"
)

var buildCmd = commandDefinitions{
	name:         "build",
	doc:          "...",
	requirements: projctx.EnvRequireFlagRecipe,
	execCommand: func(projCtx *projctx.ProjectContext) {
		operations.BuildProject(projCtx, false)
	},
}

var devCmd = commandDefinitions{
	name:         "dev",
	doc:          "...",
	requirements: projctx.EnvRequireFlagComMojang | projctx.EnvRequireFlagRecipe,
	execCommand: func(projCtx *projctx.ProjectContext) {
		operations.CopyToDev(projCtx)
	},
}

var packCmd = commandDefinitions{
	name:         "pack",
	doc:          "...",
	requirements: projctx.EnvRequireFlagRecipe,
	execCommand: func(projCtx *projctx.ProjectContext) {
		operations.PackProject(projCtx)
	},
}
