package recipe

import (
	"encoding/json"

	"github.com/redrock/autocrafter/jsonst"
	"github.com/redrock/autocrafter/recipe/formatv"
	v1 "github.com/redrock/autocrafter/recipe/versions/v1"
)

type RecipeType uint8

const (
	BehaviourPackRecipeType RecipeType = iota + 1
	ResourcePackRecipeType
	AddonRecipeType
)

type UUIDPack struct {
	Single *jsonst.UUID
	RP     *jsonst.UUID
	BP     *jsonst.UUID
}

type Recipe struct {
	/* Config */
	Type     RecipeType
	Artifact string

	/* Header */
	Version          *jsonst.SemVer
	UUIDs            UUIDPack
	MinEngineVersion *jsonst.SemVer

	/* Modules */
	Modules []Module

	/* Meta */
	Authors []string
	License string
}

func (recipe *Recipe) UnmarshalJSON(data []byte) error {
	formatVersion, jsonErr := formatv.Get(data)
	if jsonErr != nil {
		return jsonErr
	}

	switch formatVersion {
	case 1:
		return loadV1(recipe, data)
	}

	return &formatv.UnsupportedFormatVersionError{Version: formatVersion}
}

func loadV1(recipe *Recipe, data []byte) error {
	var rawRecipe v1.Recipe
	json_err := json.Unmarshal(data, &rawRecipe)
	if json_err != nil {
		return json_err
	}

	recipe.Type = RecipeType(rawRecipe.Config.Type)
	recipe.Artifact = rawRecipe.Config.Artifact

	recipe.Version = rawRecipe.Header.Version

	recipe.UUIDs.Single = rawRecipe.Header.UUID
	if rawRecipe.Header.UUIDs != nil {
		recipe.UUIDs.BP = rawRecipe.Header.UUIDs.BP
		recipe.UUIDs.RP = rawRecipe.Header.UUIDs.RP
	}

	recipe.MinEngineVersion = rawRecipe.Header.MinEngineVersion

	modLen := len(rawRecipe.Modules)
	recipe.Modules = make([]Module, modLen)
	for i, mod := range rawRecipe.Modules {
		recipe.Modules[i].Description = mod.Description
		recipe.Modules[i].Type = ModuleType(mod.Type)
		recipe.Modules[i].Version = mod.Version
		recipe.Modules[i].UUID = mod.UUID
	}

	if rawRecipe.Meta != nil {
		recipe.Authors = rawRecipe.Meta.Authors
		recipe.License = rawRecipe.Meta.License
	}

	return nil
}
