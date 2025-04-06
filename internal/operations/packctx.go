package operations

import (
	"path/filepath"

	"github.com/n-loco/mcbuild/internal/operations/internal/manifest"
	"github.com/n-loco/mcbuild/internal/projctx"
	"github.com/n-loco/mcbuild/internal/projctx/recipe"
)

type packContext struct {
	*projctx.ProjectContext
	packType    recipe.PackType
	release     bool
	scriptDeps  map[string]manifest.Dependency
	packDistDir string
}

func createPackContext(projCtx *projctx.ProjectContext, packType recipe.PackType, release bool) (packCtx packContext) {
	recipeType := projCtx.Recipe.Type

	packCtx.ProjectContext = projCtx
	packCtx.packType = packType
	packCtx.scriptDeps = make(map[string]manifest.Dependency)
	packCtx.release = release

	baseDir := filepath.Join(projCtx.DistDir, "._obj")

	if release {
		baseDir = filepath.Join(baseDir, "release")
	} else {
		baseDir = filepath.Join(baseDir, "debug")
	}

	if recipeType == recipe.RecipeTypeAddon {
		packCtx.packDistDir = filepath.Join(baseDir, packType.Abbr())
	} else {
		packCtx.packDistDir = baseDir
	}

	return
}
