package operations

import (
	"archive/zip"
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/projctx"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

func PackProject(projCtx *projctx.ProjectContext, debugP bool) {
	var diagnostic = projCtx.Diagnostic

	var projRecipe = projCtx.Recipe
	var recipeType = projRecipe.Type

	BuildProject(projCtx, !debugP)

	if diagnostic.HasErrors() {
		return
	}

	var packFilePath = filepath.Join(projCtx.DistDir, projRecipe.Artifact)

	if debugP {
		packFilePath += ".debug"
	}

	if recipeType == projfiles.RecipeTypeAddOn {
		packFilePath += ".mcaddon"
	} else {
		packFilePath += ".mcpack"
	}

	var packFile, packFileErr = os.Create(packFilePath)
	if packFileErr != nil {
		diagnostic.AppendError(alert.WrappGoError(packFileErr))
		return
	}

	var zipW = zip.NewWriter(packFile)
	defer zipW.Close()

	var tmpDir, tmpDirErr = os.MkdirTemp(os.TempDir(), "_mcbpacking")
	if tmpDirErr != nil {
		diagnostic.AppendError(alert.WrappGoError(tmpDirErr))
		return
	}
	defer os.RemoveAll(tmpDir)

	if recipeType == projfiles.RecipeTypeAddOn {
		var bpCtx, rpCtx = projCtx.AddonContext(!debugP)

		copyPackToTempDir(tmpDir, &bpCtx)
		copyPackToTempDir(tmpDir, &rpCtx)
	} else {
		var packCtx = projCtx.PackContext(!debugP)
		copyPackToTempDir(tmpDir, &packCtx)
	}

	if diagnostic.HasErrors() {
		return
	}

	var tmpFS = os.DirFS(tmpDir)

	var zipErr = zipW.AddFS(tmpFS)
	diagnostic.AppendError(alert.WrappGoError(zipErr))
}

func copyPackToTempDir(tempDir string, packCtx *projctx.PackContext) {
	var diagnostic = packCtx.Diagnostic
	var recipeType = packCtx.Recipe.Type

	var packDir = tempDir

	if recipeType == projfiles.RecipeTypeAddOn {
		packDir = filepath.Join(packDir, packCtx.PackDirName)
	}

	var err = os.CopyFS(packDir, os.DirFS(packCtx.PackDistDir))

	diagnostic.AppendError(alert.WrappGoError(err))
}
