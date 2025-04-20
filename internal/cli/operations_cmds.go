package cli

import (
	"github.com/n-loco/bpbuild/internal/operations"
	"github.com/n-loco/bpbuild/internal/projctx"
)

var buildCmd = simpleCommand{
	commandInfo: commandInfo{
		name: "build",
		doc:  "...",
	},
	requirements: projctx.EnvRequireFlagRecipe,
	execCommand: func(projCtx *projctx.ProjectContext) {
		operations.BuildProject(projCtx, false)
	},
}

var devCmd = simpleCommand{
	commandInfo: commandInfo{
		name: "dev",
		doc:  "...",
	},
	requirements: projctx.EnvRequireFlagComMojang | projctx.EnvRequireFlagRecipe,
	execCommand: func(projCtx *projctx.ProjectContext) {
		operations.CopyToDev(projCtx)
	},
}

var packCmd = simpleCommand{
	commandInfo: commandInfo{
		name: "pack",
		doc:  "...",
	},
	requirements: projctx.EnvRequireFlagRecipe,
	execCommand: func(projCtx *projctx.ProjectContext) {
		operations.PackProject(projCtx)
	},
}
