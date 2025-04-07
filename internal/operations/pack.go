package operations

import (
	"archive/zip"
	"os"
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

func PackProject(projCtx *projctx.ProjectContext) {
	projRecipe := projCtx.Recipe
	recipeType := projRecipe.Type

	BuildProject(projCtx, true)

	packFilePath := filepath.Join(projCtx.DistDir, projRecipe.Artifact)

	if recipeType == recipe.RecipeTypeAddon {
		packFilePath += ".mcaddon"
	} else {
		packFilePath += ".mcpack"
	}

	packFile, cFErr := os.Create(packFilePath)
	if cFErr != nil {
		panic("TODO EERH CREATING PACK: " + cFErr.Error())
	}

	zipW := zip.NewWriter(packFile)
	defer zipW.Close()

	tmpDir, tdErr := os.MkdirTemp(os.TempDir(), "_mcbpacking")
	if tdErr != nil {
		panic("TODO EERH CREATING PACK: " + tdErr.Error())
	}
	defer os.RemoveAll(tmpDir)

	if recipeType == recipe.RecipeTypeAddon {
		bpCtx := createPackContext(projCtx, recipe.PackTypeBehavior, true)
		copyPackToTempDir(tmpDir, &bpCtx)

		rpCtx := createPackContext(projCtx, recipe.PackTypeResource, true)
		copyPackToTempDir(tmpDir, &rpCtx)
	} else {
		packCtx := createPackContext(projCtx, recipeType.PackType(), true)
		copyPackToTempDir(tmpDir, &packCtx)
	}

	tmpFS := os.DirFS(tmpDir)

	zipW.AddFS(tmpFS)
}

func copyPackToTempDir(tempDir string, packCtx *packContext) {
	recipeType := packCtx.Recipe.Type

	packDir := tempDir

	if recipeType == recipe.RecipeTypeAddon {
		packDir = filepath.Join(packDir, packCtx.packDirName)
	}

	os.CopyFS(packDir, os.DirFS(packCtx.packDistDir))
}
