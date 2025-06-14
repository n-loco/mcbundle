package projctx

import (
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/projfiles"
)

func (packCtx *PackContext) ModuleContext(recipeModule *projfiles.Module) (modCtx ModuleContext) {
	modCtx.PackContext = packCtx
	modCtx.RecipeModule = recipeModule
	modCtx.ModSourcePath = filepath.Join(packCtx.SourceDir, recipeModule.Type.String())
	return
}

type ModuleContext struct {
	*PackContext
	RecipeModule  *projfiles.Module
	ModSourcePath string
}
