package cli

import (
	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/operations"
	"github.com/mcbundle/mcbundle/internal/projctx"
)

func buildCmd(argv *argvIterator, diagnostic alert.Diagnostic) {
	var recipe = projctx.CreateProjectContext(projctx.EnvRequireFlagRecipe, diagnostic)

	var isRelease = false

	if argv.hasNext() {
		var arg = argv.consume()

		if arg == "--release" {
			isRelease = true
		} else {
			diagnostic.AppendWarning(alert.AlertF("unknown option: %s", arg))
		}
	}

	if diagnostic.HasErrors() {
		return
	}

	operations.BuildProject(&recipe, isRelease)
}

func devCmd(argv *argvIterator, diagnostic alert.Diagnostic) {
	var recipe = projctx.CreateProjectContext(projctx.EnvRequireFlagRecipe|projctx.EnvRequireFlagComMojang, diagnostic)

	if diagnostic.HasErrors() {
		return
	}

	operations.CopyToDev(&recipe)
}

func packCmd(argv *argvIterator, diagnostic alert.Diagnostic) {
	var recipe = projctx.CreateProjectContext(projctx.EnvRequireFlagRecipe, diagnostic)

	var isDebug = false

	if argv.hasNext() {
		var arg = argv.consume()

		if arg == "--debug" {
			isDebug = true
		} else {
			diagnostic.AppendWarning(alert.AlertF("unknown option: %s", arg))
		}
	}

	if diagnostic.HasErrors() {
		return
	}

	operations.PackProject(&recipe, isDebug)
}
