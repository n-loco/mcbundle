package projctx

import (
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/projctx/recipe"
)

func (packCtx *PackContext) ModuleContext(recipeModule *recipe.Module) (modCtx ModuleContext) {
	modCtx.PackContext = packCtx
	modCtx.RecipeModule = recipeModule
	modCtx.ModSourcePath = filepath.Join(packCtx.SourceDir, recipeModule.Type.String())
	return
}

type ModuleContext struct {
	*PackContext
	RecipeModule  *recipe.Module
	ModSourcePath string
}
