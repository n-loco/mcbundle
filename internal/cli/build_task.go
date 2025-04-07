package cli

import (
	"github.com/n-loco/bpbuild/internal/operations"
	"github.com/n-loco/bpbuild/internal/projctx"
)

var buildTask = TaskDefs{
	Requires: projctx.EnvRequireFlagRecipe,
	Name:     "build",
	Doc:      "bundles JS/TS and copies content (e. g: data or resources) files into the dist directory.",
	Execute: func(projctx *projctx.ProjectContext) {
		operations.BuildProject(projctx, false)
	},
}
