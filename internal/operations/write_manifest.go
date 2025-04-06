package operations

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/n-loco/mcbuild/internal/operations/internal/manifest"
	"github.com/n-loco/mcbuild/internal/projctx/recipe"
)

func writeManifest(packCtx *packContext, builtModules []manifest.Module, foundDeps []manifest.Dependency) {
	manifest := createManifest(packCtx)

	manifest.Modules = builtModules
	manifest.Dependencies = append(manifest.Dependencies, foundDeps...)

	manifestPath := filepath.Join(packCtx.packDistDir, "manifest.json")
	manifestData, _ := json.MarshalIndent(&manifest, "", "  ")

	os.WriteFile(manifestPath, manifestData, os.ModePerm)
}

func createManifest(packCtx *packContext) (mcManifest *manifest.Manifest) {
	recipeType := packCtx.Recipe.Type

	mcManifest = new(manifest.Manifest)
	mcManifest.FormatVersion = 2

	setHeader(&mcManifest.Header, packCtx)

	if recipeType == recipe.RecipeTypeAddon {
		mcManifest.Dependencies = append(mcManifest.Dependencies, createAddonDependency(packCtx))
	}

	setMeta(&mcManifest.Meta, packCtx)

	return
}

func setHeader(header *manifest.Header, packCtx *packContext) {
	projRecipe := packCtx.Recipe
	packType := packCtx.packType

	if projRecipe.Type == recipe.RecipeTypeAddon {
		switch packType {
		case recipe.PackTypeBehaviour:
			header.UUID = projRecipe.UUIDs.BP
		case recipe.PackTypeResource:
			header.UUID = projRecipe.UUIDs.RP
			header.PackScope = manifest.PackScopeWorld
		default:
			panic("💀")
		}
	} else {
		header.UUID = packCtx.Recipe.UUIDs.Single
	}

	header.Description = "pack.description"
	header.Name = "pack.name"
	header.Version = projRecipe.Version
	header.MinEngineVersion = projRecipe.MinEngineVersion
}

func createAddonDependency(packCtx *packContext) (dependency manifest.Dependency) {
	projRecipe := packCtx.Recipe
	packType := packCtx.packType

	switch packType {
	case recipe.PackTypeBehaviour:
		dependency.UUID = projRecipe.UUIDs.RP
	case recipe.PackTypeResource:
		dependency.UUID = projRecipe.UUIDs.BP
	default:
		panic("💀")
	}

	dependency.Version = projRecipe.Version

	return
}

func setMeta(meta *manifest.Meta, packCtx *packContext) {
	projRecipe := packCtx.Recipe

	meta.Authors = projRecipe.Authors
	meta.License = projRecipe.License
}
