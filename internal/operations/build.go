package operations

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/operations/internal/manifest"
	"github.com/mcbundle/mcbundle/internal/projctx"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

func BuildProject(projCtx *projctx.ProjectContext, release bool) (diagnostic *alert.Diagnostic) {
	projType := projCtx.Recipe.Type

	if projType == projfiles.RecipeTypeAddOn {
		bpCtx, rpCtx := projCtx.AddonContext(release)

		diagnostic = diagnostic.Append(buildPack(&bpCtx))
		diagnostic = diagnostic.Append(buildPack(&rpCtx))
	} else {
		packCtx := projCtx.PackContext(release)
		diagnostic = diagnostic.Append(buildPack(&packCtx))
	}

	return
}

func buildPack(packCtx *projctx.PackContext) (diagnostic *alert.Diagnostic) {
	projRecipe := packCtx.Recipe
	packType := packCtx.PackType

	buildPath := packCtx.PackDistDir
	if _, err := os.Stat(buildPath); err == nil {
		os.RemoveAll(buildPath)
	}

	var foundDeps []manifest.Dependency
	var builtModules []manifest.Module

	for _, recipeModule := range projRecipe.Modules {
		if recipeModule.BelongsTo() != packType {
			continue
		}

		modCtx := packCtx.ModuleContext(&recipeModule)

		mod, buildModDiag := buildModule(&modCtx)

		diagnostic = diagnostic.Append(buildModDiag)

		builtModules = append(builtModules, mod)
	}

	if diagnostic.HasErrors() {
		return
	}

	for _, dep := range packCtx.ScriptDependencies() {
		foundDeps = append(foundDeps, manifest.Dependency{
			ModuleName: dep.Name,
			Version:    dep.Version,
		})
	}

	var packIconPath = filepath.Join(packCtx.PackDistDir, "pack_icon.png")
	packIconPath, _ = filepath.Rel(packCtx.WorkDir, packIconPath)

	var packIconFile, packIconErr = os.Create(packIconPath)
	if packIconErr != nil {
		diagnostic.AppendError(alert.WrappGoError(packIconErr))
	} else {
		packIconFile.Write(packCtx.PackIcon)
		packIconFile.Close()
	}

	writeManifest(packCtx, builtModules, foundDeps)

	return
}

func buildModule(modCtx *projctx.ModuleContext) (mod manifest.Module, diagnostic *alert.Diagnostic) {
	recipeModule := modCtx.RecipeModule

	switch recipeModule.Type {
	case projfiles.ModuleTypeData:
		fallthrough
	case projfiles.ModuleTypeResources:
		{
			diagnostic = diagnostic.Append(copyDataToBuild(modCtx.ModSourcePath, modCtx.PackDistDir))
		}
	case projfiles.ModuleTypeServer:
		{
			diagnostic = diagnostic.Append(esbuild(modCtx))
			if !diagnostic.HasErrors() {
				mod.Entry = "scripts/server.js"
				mod.Language = "javascript"
			}
		}
	default:
		panic("invalid module")
	}

	if diagnostic.HasErrors() {
		return
	}

	mod.UUID = recipeModule.UUID
	mod.Version = recipeModule.Version
	mod.Type = manifest.ModuleTypeFromRecipeModuleType(recipeModule.Type)

	return
}

func copyDataToBuild(from string, to string) (diagnostic *alert.Diagnostic) {
	return diagnostic.AppendError(alert.WrappGoError(os.CopyFS(to, os.DirFS(from))))
}
