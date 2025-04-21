package projctx

import (
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

func (packCtx *PackContext) ModuleContext(recipeModule *recipe.RecipeModule) (modCtx ModuleContext) {
	recipeModType := recipeModule.Type

	modCtx.PackContext = packCtx
	modCtx.RecipeModule = recipeModule

	modCtx.ModSourcePath = filepath.Join(packCtx.SourceDir, recipeModType.String())
	return
}

type ModuleContext struct {
	*PackContext
	RecipeModule  *recipe.RecipeModule
	ModSourcePath string
}
