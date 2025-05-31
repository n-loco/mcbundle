package cli

import (
	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/operations"
	"github.com/mcbundle/mcbundle/internal/projctx"
)

type buildOptsObj struct {
	release bool
}

type packOptsObj struct {
	debug bool
}

var buildCmd = createOperationCommand(
	commandInfo{
		name: "build",
		doc:  "...",
	},
	projctx.EnvRequireFlagRecipe,
	func(obj *buildOptsObj, projCtx *projctx.ProjectContext) (diagnostic *alert.Diagnostic) {
		diagnostic = diagnostic.Append(operations.BuildProject(projCtx, obj.release))
		return
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
	func(obj *empty, projCtx *projctx.ProjectContext) (diagnostic *alert.Diagnostic) {
		diagnostic = diagnostic.Append(operations.CopyToDev(projCtx))
		return
	},
	nil,
)

var packCmd = createOperationCommand(
	commandInfo{
		name: "pack",
		doc:  "...",
	},
	projctx.EnvRequireFlagRecipe,
	func(obj *packOptsObj, projCtx *projctx.ProjectContext) (diagnostic *alert.Diagnostic) {
		diagnostic = diagnostic.Append(operations.PackProject(projCtx, obj.debug))
		return
	},
	[]*operationOption[packOptsObj]{
		{
			optionInfo: optionInfo{
				name: "--debug",
			},
			process: func(o *packOptsObj, optSlice []string) int {
				o.debug = true
				return 0
			},
		},
	},
)
