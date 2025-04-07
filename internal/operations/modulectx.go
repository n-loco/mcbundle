package operations

import (
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

type moduleContext struct {
	*packContext
	recipeModule  *recipe.RecipeModule
	modSourcePath string
}

func createModuleContext(packCtx *packContext, recipeModule *recipe.RecipeModule) (modCtx moduleContext) {
	recipeModType := recipeModule.Type

	modCtx.packContext = packCtx
	modCtx.recipeModule = recipeModule

	modCtx.modSourcePath = filepath.Join(packCtx.SourceDir, recipeModType.String())

	return
}
