package operations

import (
	"archive/zip"
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/projctx"
	"github.com/mcbundle/mcbundle/internal/projctx/recipe"
)

func PackProject(projCtx *projctx.ProjectContext, debugP bool) (diagnostic *alert.Diagnostic) {
	projRecipe := projCtx.Recipe
	recipeType := projRecipe.Type

	diagnostic = diagnostic.Append(BuildProject(projCtx, !debugP))

	if diagnostic.HasErrors() {
		return
	}

	packFilePath := filepath.Join(projCtx.DistDir, projRecipe.Artifact)

	if debugP {
		packFilePath += ".debug"
	}

	if recipeType == recipe.RecipeTypeAddOn {
		packFilePath += ".mcaddon"
	} else {
		packFilePath += ".mcpack"
	}

	packFile, packFileErr := os.Create(packFilePath)
	if packFileErr != nil {
		diagnostic = diagnostic.AppendError(alert.WrappGoError(packFileErr))
		return
	}

	zipW := zip.NewWriter(packFile)
	defer zipW.Close()

	tmpDir, tmpDirErr := os.MkdirTemp(os.TempDir(), "_mcbpacking")
	if tmpDirErr != nil {
		diagnostic = diagnostic.AppendError(alert.WrappGoError(tmpDirErr))
		return
	}
	defer os.RemoveAll(tmpDir)

	if recipeType == recipe.RecipeTypeAddOn {
		bpCtx, rpCtx := projCtx.AddonContext(!debugP)

		diagnostic = diagnostic.Append(copyPackToTempDir(tmpDir, &bpCtx))
		diagnostic = diagnostic.Append(copyPackToTempDir(tmpDir, &rpCtx))
	} else {
		packCtx := projCtx.PackContext(!debugP)
		diagnostic = diagnostic.Append(copyPackToTempDir(tmpDir, &packCtx))
	}

	if diagnostic.HasErrors() {
		return
	}

	tmpFS := os.DirFS(tmpDir)

	zipErr := zipW.AddFS(tmpFS)
	diagnostic = diagnostic.AppendError(alert.WrappGoError(zipErr))

	return
}

func copyPackToTempDir(tempDir string, packCtx *projctx.PackContext) (diagnostic *alert.Diagnostic) {
	recipeType := packCtx.Recipe.Type

	packDir := tempDir

	if recipeType == recipe.RecipeTypeAddOn {
		packDir = filepath.Join(packDir, packCtx.PackDirName)
	}

	err := os.CopyFS(packDir, os.DirFS(packCtx.PackDistDir))

	diagnostic = diagnostic.AppendError(alert.WrappGoError(err))

	return
}
