package operations

import (
	"os"

	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

func CopyToDev(projCtx *projctx.ProjectContext) {
	recipeType := projCtx.Recipe.Type

	BuildProject(projCtx, false)

	if recipeType == recipe.RecipeTypeAddon {
		bpCtx := createPackContext(projCtx, recipe.PackTypeBehavior, false)
		copyPackToDev(&bpCtx)

		rpCtx := createPackContext(projCtx, recipe.PackTypeResource, false)
		copyPackToDev(&rpCtx)
	} else {
		packCtx := createPackContext(projCtx, recipeType.PackType(), false)
		copyPackToDev(&packCtx)
	}
}

func copyPackToDev(packCtx *packContext) {
	os.RemoveAll(packCtx.packDevDir)
	err := os.CopyFS(packCtx.packDevDir, os.DirFS(packCtx.packDistDir))

	if err != nil {
		panic("TODO ERR HANDLING: " + err.Error())
	}
}
