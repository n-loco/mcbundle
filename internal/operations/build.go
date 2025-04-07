package operations

import (
	"log"
	"os"

	"github.com/n-loco/bpbuild/internal/operations/internal/manifest"
	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/projctx/recipe"
)

func BuildProject(projCtx *projctx.ProjectContext, release bool) {
	projType := projCtx.Recipe.Type

	if projType == recipe.RecipeTypeAddon {
		bpCtx := createPackContext(projCtx, recipe.PackTypeBehavior, release)
		buildPack(&bpCtx)

		rpCtx := createPackContext(projCtx, recipe.PackTypeResource, release)
		buildPack(&rpCtx)
	} else {
		packCtx := createPackContext(projCtx, projType.PackType(), release)
		buildPack(&packCtx)
	}
}

func buildPack(packCtx *packContext) {
	projRecipe := packCtx.Recipe
	packType := packCtx.packType

	buildPath := packCtx.packDistDir
	if _, err := os.Stat(buildPath); err == nil {
		os.RemoveAll(buildPath)
	}

	var foundDeps []manifest.Dependency
	var builtModules []manifest.Module

	for _, recipeModule := range projRecipe.Modules {
		if recipeModule.Type.PackType() != packType {
			continue
		}

		modCtx := createModuleContext(packCtx, &recipeModule)

		mod, err := buildModule(&modCtx)
		if err != nil {
			log.Print("TODO! error handling buildModule: " + err.Error())
		}

		builtModules = append(builtModules, mod)
	}

	for _, dep := range packCtx.scriptDeps {
		foundDeps = append(foundDeps, dep)
	}

	writeManifest(packCtx, builtModules, foundDeps)
}

func buildModule(modCtx *moduleContext) (mod manifest.Module, err error) {
	recipeModule := modCtx.recipeModule

	switch recipeModule.Type {
	case recipe.RecipeModuleTypeData:
		fallthrough
	case recipe.RecipeModuleTypeResources:
		{
			err = copyDataToBuild(modCtx.modSourcePath, modCtx.packDistDir)
		}
	case recipe.RecipeModuleTypeServer:
		{
			err = esbuild(modCtx)
			if err == nil {
				mod.Entry = "scripts/server.js"
				mod.Language = "javascript"
			}
		}
	default:
		panic("invalid module")
	}

	if err == nil {
		mod.Description = recipeModule.Description
		mod.UUID = recipeModule.UUID
		mod.Version = recipeModule.Version
		mod.Type = manifest.ModuleTypeFromRecipeModuleType(recipeModule.Type)
	}

	return
}

func copyDataToBuild(from string, to string) error {
	return os.CopyFS(to, os.DirFS(from))
}
