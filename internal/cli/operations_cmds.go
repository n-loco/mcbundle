package cli

import (
	"github.com/n-loco/bpbuild/internal/operations"
	"github.com/n-loco/bpbuild/internal/projctx"
)

type buildOptsObj struct {
	release bool
}

type packOptsObj struct {
}

var buildCmd = createOperationCommand(
	commandInfo{
		name: "build",
		doc:  "...",
	},
	projctx.EnvRequireFlagRecipe,
	func(obj *buildOptsObj, projCtx *projctx.ProjectContext) {
		operations.BuildProject(projCtx, obj.release)
	},
	[]*operationOption[buildOptsObj]{
		{
			optionInfo: optionInfo{
				name: "--release",
			},
			process: func(o *buildOptsObj, optSlice []string) int {
				o.release = true
				return 0
			},
		},
	},
)

var devCmd = createOperationCommand(
	commandInfo{
		name: "dev",
		doc:  "...",
	},
	projctx.EnvRequireFlagRecipe|projctx.EnvRequireFlagComMojang,
	func(obj *empty, projCtx *projctx.ProjectContext) {
		operations.CopyToDev(projCtx)
	},
	nil,
)

var packCmd = createOperationCommand(
	commandInfo{
		name: "pack",
		doc:  "...",
	},
	projctx.EnvRequireFlagRecipe,
	func(obj *packOptsObj, projCtx *projctx.ProjectContext) {
		operations.PackProject(projCtx)
	},
	nil,
)
