package operations

import (
	"archive/zip"
	"os"
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/alert"
	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

func PackProject(projCtx *projctx.ProjectContext) (diagnostic *alert.Diagnostic) {
	projRecipe := projCtx.Recipe
	recipeType := projRecipe.Type

	diagnostic = diagnostic.Append(BuildProject(projCtx, true))

	if diagnostic.HasErrors() {
		return
	}

	packFilePath := filepath.Join(projCtx.DistDir, projRecipe.Artifact)

	if recipeType == recipe.RecipeTypeAddon {
		packFilePath += ".mcaddon"
	} else {
		packFilePath += ".mcpack"
	}

	packFile, packFileErr := os.Create(packFilePath)
	if packFileErr != nil {
		diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(packFileErr))
		return
	}

	zipW := zip.NewWriter(packFile)
	defer zipW.Close()

	tmpDir, tmpDirErr := os.MkdirTemp(os.TempDir(), "_mcbpacking")
	if tmpDirErr != nil {
		diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(tmpDirErr))
		return
	}
	defer os.RemoveAll(tmpDir)

	if recipeType == recipe.RecipeTypeAddon {
		bpCtx := createPackContext(projCtx, recipe.PackTypeBehavior, true)
		diagnostic = diagnostic.Append(copyPackToTempDir(tmpDir, &bpCtx))

		rpCtx := createPackContext(projCtx, recipe.PackTypeResource, true)
		diagnostic = diagnostic.Append(copyPackToTempDir(tmpDir, &rpCtx))
	} else {
		packCtx := createPackContext(projCtx, recipeType.PackType(), true)
		diagnostic = diagnostic.Append(copyPackToTempDir(tmpDir, &packCtx))
	}

	if diagnostic.HasErrors() {
		return
	}

	tmpFS := os.DirFS(tmpDir)

	zipErr := zipW.AddFS(tmpFS)
	diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(zipErr))

	return
}

func copyPackToTempDir(tempDir string, packCtx *packContext) (diagnostic *alert.Diagnostic) {
	recipeType := packCtx.Recipe.Type

	packDir := tempDir

	if recipeType == recipe.RecipeTypeAddon {
		packDir = filepath.Join(packDir, packCtx.packDirName)
	}

	err := os.CopyFS(packDir, os.DirFS(packCtx.packDistDir))

	diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(err))

	return
}
