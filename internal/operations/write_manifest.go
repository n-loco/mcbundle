package operations

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/operations/internal/manifest"
	"github.com/mcbundle/mcbundle/internal/projctx"
	"github.com/mcbundle/mcbundle/internal/projctx/recipe"
)

func writeManifest(packCtx *projctx.PackContext, builtModules []manifest.Module, foundDeps []manifest.Dependency) {
	manifest := createManifest(packCtx)

	manifest.Modules = builtModules
	manifest.Dependencies = append(manifest.Dependencies, foundDeps...)

	manifestPath := filepath.Join(packCtx.PackDistDir, "manifest.json")
	manifestData, _ := json.MarshalIndent(&manifest, "", "  ")

	os.WriteFile(manifestPath, manifestData, os.ModePerm)
}

func createManifest(packCtx *projctx.PackContext) (mcManifest *manifest.Manifest) {
	recipeType := packCtx.Recipe.Type

	mcManifest = new(manifest.Manifest)
	mcManifest.FormatVersion = 2

	setHeader(&mcManifest.Header, packCtx)

	if recipeType == recipe.RecipeTypeAddOn {
		mcManifest.Dependencies = append(mcManifest.Dependencies, createAddonDependency(packCtx))
	}

	setMeta(&mcManifest.MetaData, packCtx)

	return
}

func setHeader(header *manifest.Header, packCtx *projctx.PackContext) {
	projRecipe := packCtx.Recipe
	packType := packCtx.PackType

	if projRecipe.Type == recipe.RecipeTypeAddOn {
		switch packType {
		case recipe.PackTypeBehavior:
			header.UUID = projRecipe.UUIDs[0]
		case recipe.PackTypeResources:
			header.UUID = projRecipe.UUIDs[1]
			header.PackScope = manifest.PackScopeWorld
		default:
			panic("💀")
		}
	} else {
		header.UUID = packCtx.Recipe.UUID
	}

	header.Name = projRecipe.Name
	header.Version = projRecipe.Version
	header.MinEngineVersion = [3]uint8{1, 21, 0}
}

func createAddonDependency(packCtx *projctx.PackContext) (dependency manifest.Dependency) {
	projRecipe := packCtx.Recipe
	packType := packCtx.PackType

	switch packType {
	case recipe.PackTypeBehavior:
		dependency.UUID = projRecipe.UUIDs[1]
	case recipe.PackTypeResources:
		dependency.UUID = projRecipe.UUIDs[0]
	default:
		panic("💀")
	}

	dependency.Version = projRecipe.Version

	return
}

func setMeta(meta *manifest.MetaData, packCtx *projctx.PackContext) {
	projRecipe := packCtx.Recipe

	meta.Authors = projRecipe.Authors
	meta.License = projRecipe.License
}
