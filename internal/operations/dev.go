package operations

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/projctx"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

func CopyToDev(projCtx *projctx.ProjectContext) (diagnostic *alert.Diagnostic) {
	recipeType := projCtx.Recipe.Type

	diagnostic = diagnostic.Append(BuildProject(projCtx, false))

	if diagnostic.HasErrors() {
		return
	}

	if recipeType == projfiles.RecipeTypeAddOn {
		bpCtx, rpCtx := projCtx.AddonContext(false)

		diagnostic = diagnostic.Append(copyPackToDev(&bpCtx))
		diagnostic = diagnostic.Append(copyPackToDev(&rpCtx))
	} else {
		packCtx := projCtx.PackContext(false)
		diagnostic = diagnostic.Append(copyPackToDev(&packCtx))
	}

	return
}

func copyPackToDev(packCtx *projctx.PackContext) (diagnostic *alert.Diagnostic) {
	os.RemoveAll(packCtx.PackDevDir)
	err := os.CopyFS(packCtx.PackDevDir, os.DirFS(packCtx.PackDistDir))

	diagnostic = diagnostic.AppendError(alert.WrappGoError(err))

	return
}
