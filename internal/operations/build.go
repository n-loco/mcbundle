package operations

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/n-loco/mcbuild/internal/jsonst"
	"github.com/n-loco/mcbuild/internal/mcmanifest"
	"github.com/n-loco/mcbuild/internal/rcontext"
	"github.com/n-loco/mcbuild/internal/rcontext/recipe"
)

type buildModuleContext struct {
	buildPath           string
	sourcePath          string
	scriptDependencyMap map[string]mcmanifest.Dependency
	recipeModule        *recipe.RecipeModule
}

func BuildProject(projectRecipe *recipe.Recipe, release bool) {
	if projectRecipe.Type == recipe.RecipeTypeAddon {
		bpCtx := rcontext.Context{
			Recipe:   projectRecipe,
			Release:  release,
			PackType: recipe.PackTypeBehaviour,
		}
		buildPack(&bpCtx)

		rpCtx := rcontext.Context{
			Recipe:   projectRecipe,
			Release:  release,
			PackType: recipe.PackTypeResource,
		}
		buildPack(&rpCtx)
	} else {
		ctx := rcontext.Context{
			Recipe:   projectRecipe,
			Release:  release,
			PackType: projectRecipe.Type.PackType(),
		}
		buildPack(&ctx)
	}
}

func buildPack(ctx *rcontext.Context) {
	buildPath := buildPath(ctx)
	if _, err := os.Stat(buildPath); err == nil {
		os.RemoveAll(buildPath)
	}

	scriptDeps := make(map[string]mcmanifest.Dependency)
	manifest := mcmanifest.CreateManifest(ctx)

	if ctx.Recipe.Type == recipe.RecipeTypeAddon {
		var uuid *jsonst.UUID

		switch ctx.PackType {
		case recipe.PackTypeBehaviour:
			uuid = ctx.Recipe.UUIDs.RP
		case recipe.PackTypeResource:
			uuid = ctx.Recipe.UUIDs.BP
		}

		manifest.Dependencies = append(manifest.Dependencies, mcmanifest.Dependency{
			Version: ctx.Recipe.Version,
			UUID:    uuid,
		})
	}

	for _, recipeModule := range ctx.Recipe.Modules {
		if recipeModule.Type.PackType() != ctx.PackType {
			continue
		}

		buildModCtx := buildModuleContext{
			buildPath:           buildPath,
			sourcePath:          moduleSourcePath(&recipeModule),
			scriptDependencyMap: scriptDeps,
			recipeModule:        &recipeModule,
		}

		mod, err := buildModule(ctx, &buildModCtx)
		if err != nil {
			log.Print("TODO! error handling buildModule")
		}

		manifest.Modules = append(manifest.Modules, mod)
	}

	for _, dep := range scriptDeps {
		manifest.Dependencies = append(manifest.Dependencies, dep)
	}

	manifestPath := filepath.Join(buildPath, "manifest.json")
	manifestData, _ := json.MarshalIndent(&manifest, "", "  ")
	os.WriteFile(manifestPath, manifestData, os.ModePerm)
}

func buildModule(ctx *rcontext.Context, buildModCtx *buildModuleContext) (mod mcmanifest.Module, err error) {
	switch buildModCtx.recipeModule.Type {
	case recipe.RecipeModuleTypeData:
		fallthrough
	case recipe.RecipeModuleTypeResources:
		{
			err = copyDataToBuild(buildModCtx.sourcePath, buildModCtx.buildPath)
		}
	case recipe.RecipeModuleTypeServer:
		{
			err = esbuild(ctx, buildModCtx)
			if err == nil {
				mod.Entry = "scripts/server.js"
				mod.Language = "javascript"
			}
		}
	default:
		panic("invalid module")
	}

	if err == nil {
		mod.Description = buildModCtx.recipeModule.Description
		mod.UUID = buildModCtx.recipeModule.UUID
		mod.Version = buildModCtx.recipeModule.Version
		mod.Type = mcmanifest.ModuleTypeFromRecipeModuleType(buildModCtx.recipeModule.Type)
	}

	return
}

func copyDataToBuild(from string, to string) error {
	return os.CopyFS(to, os.DirFS(from))
}
