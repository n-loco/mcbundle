package operations

import (
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/operations/internal/manifest"
	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

type packContext struct {
	*projctx.ProjectContext
	packType    recipe.PackType
	release     bool
	scriptDeps  map[string]manifest.Dependency
	packDistDir string
	packDirName string
	packDevDir  string
}

func createPackContext(projCtx *projctx.ProjectContext, packType recipe.PackType, release bool) (packCtx packContext) {
	projRecipe := projCtx.Recipe
	recipeType := projRecipe.Type

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

	packCtx.packDirName = projRecipe.MojangID

	if recipeType == recipe.RecipeTypeAddon {
		packCtx.packDirName += "_" + packType.Abbr()
	}

	if len(projCtx.ComMojangDir) > 0 {
		packCtx.packDevDir = filepath.Join(packCtx.ComMojangDir, "development_"+packType.ComMojangID(), packCtx.packDirName)
	}

	return
}
