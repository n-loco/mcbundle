package operations

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/projctx"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

func CopyToDev(projCtx *projctx.ProjectContext) {
	var diagnostic = projCtx.Diagnostic

	var recipeType = projCtx.Recipe.Type

	BuildProject(projCtx, false)

	if diagnostic.HasErrors() {
		return
	}

	if recipeType == projfiles.RecipeTypeAddOn {
		var bpCtx, rpCtx = projCtx.AddonContext(false)

		copyPackToDev(&bpCtx)
		copyPackToDev(&rpCtx)
	} else {
		var packCtx = projCtx.PackContext(false)
		copyPackToDev(&packCtx)
	}
}

func copyPackToDev(packCtx *projctx.PackContext) {
	var diagnostic = packCtx.Diagnostic

	os.RemoveAll(packCtx.PackDevDir)
	var err = os.CopyFS(packCtx.PackDevDir, os.DirFS(packCtx.PackDistDir))

	diagnostic.AppendError(alert.WrappGoError(err))
}
