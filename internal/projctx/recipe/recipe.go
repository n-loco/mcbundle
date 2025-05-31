package recipe

import (
	"encoding/json"

	"github.com/mcbundle/mcbundle/internal/jsonst"
	"github.com/mcbundle/mcbundle/internal/projctx/recipe/internal/formatv"
	v1 "github.com/mcbundle/mcbundle/internal/projctx/recipe/internal/versions/v1"
)

type UUIDPack struct {
	Single *jsonst.UUID
	RP     *jsonst.UUID
	BP     *jsonst.UUID
}

type Recipe struct {
	/* Config */
	Type     RecipeType
	MojangID string
	Artifact string

	/* Header */
	Version          *jsonst.SemVer
	UUIDs            UUIDPack
	MinEngineVersion [3]uint8

	/* Modules */
	Modules []RecipeModule

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
	recipe.MojangID = rawRecipe.Config.MojangID
	recipe.Artifact = rawRecipe.Config.Artifact

	recipe.Version = rawRecipe.Header.Version

	recipe.UUIDs.Single = rawRecipe.Header.UUID
	if rawRecipe.Header.UUIDs != nil {
		recipe.UUIDs.BP = rawRecipe.Header.UUIDs.BP
		recipe.UUIDs.RP = rawRecipe.Header.UUIDs.RP
	}

	recipe.MinEngineVersion = rawRecipe.Header.MinEngineVersion

	modLen := len(rawRecipe.Modules)
	recipe.Modules = make([]RecipeModule, modLen)
	for i, mod := range rawRecipe.Modules {
		recipe.Modules[i].Description = mod.Description
		recipe.Modules[i].Type = RecipeModuleType(mod.Type)
		recipe.Modules[i].Version = mod.Version
		recipe.Modules[i].UUID = mod.UUID
	}

	if rawRecipe.Meta != nil {
		recipe.Authors = rawRecipe.Meta.Authors
		recipe.License = rawRecipe.Meta.License
	}

	return nil
}
