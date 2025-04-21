package operations

import (
	"os"
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/alert"
	"github.com/n-loco/bpbuild/internal/assets"
	"github.com/n-loco/bpbuild/internal/operations/internal/manifest"
	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

func BuildProject(projCtx *projctx.ProjectContext, release bool) (diagnostic *alert.Diagnostic) {
	projType := projCtx.Recipe.Type

	if projType == recipe.RecipeTypeAddon {
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
		if recipeModule.Type.PackType() != packType {
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

	packIconFile, _ := os.Create(filepath.Join(packCtx.PackDistDir, "pack_icon.png"))
	defer packIconFile.Close()

	packIconFile.Write(assets.DefaultPackIcon)

	writeManifest(packCtx, builtModules, foundDeps)

	return
}

func buildModule(modCtx *projctx.ModuleContext) (mod manifest.Module, diagnostic *alert.Diagnostic) {
	recipeModule := modCtx.RecipeModule

	switch recipeModule.Type {
	case recipe.RecipeModuleTypeData:
		fallthrough
	case recipe.RecipeModuleTypeResources:
		{
			diagnostic = diagnostic.Append(copyDataToBuild(modCtx.ModSourcePath, modCtx.PackDistDir))
		}
	case recipe.RecipeModuleTypeServer:
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

	mod.Description = recipeModule.Description
	mod.UUID = recipeModule.UUID
	mod.Version = recipeModule.Version
	mod.Type = manifest.ModuleTypeFromRecipeModuleType(recipeModule.Type)

	return
}

func copyDataToBuild(from string, to string) (diagnostic *alert.Diagnostic) {
	return diagnostic.AppendError(alert.NewGoErrWrapperAlert(os.CopyFS(to, os.DirFS(from))))
}
