package operations

import (
	"os"

	"github.com/n-loco/bpbuild/internal/alert"
	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

func CopyToDev(projCtx *projctx.ProjectContext) (diagnostic *alert.Diagnostic) {
	recipeType := projCtx.Recipe.Type

	diagnostic = diagnostic.Append(BuildProject(projCtx, false))

	if diagnostic.HasErrors() {
		return
	}

	if recipeType == recipe.RecipeTypeAddon {
		bpCtx := createPackContext(projCtx, recipe.PackTypeBehavior, false)
		diagnostic = diagnostic.Append(copyPackToDev(&bpCtx))

		rpCtx := createPackContext(projCtx, recipe.PackTypeResource, false)
		diagnostic = diagnostic.Append(copyPackToDev(&rpCtx))
	} else {
		packCtx := createPackContext(projCtx, recipeType.PackType(), false)
		diagnostic = diagnostic.Append(copyPackToDev(&packCtx))
	}

	return
}

func copyPackToDev(packCtx *packContext) (diagnostic *alert.Diagnostic) {
	os.RemoveAll(packCtx.packDevDir)
	err := os.CopyFS(packCtx.packDevDir, os.DirFS(packCtx.packDistDir))

	diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(err))

	return
}
