package recipe

import (
	"encoding/json"
	"fmt"

	"github.com/redrock/autocrafter/recipe/getformatv"
	v1 "github.com/redrock/autocrafter/recipe/versions/v1"
	"github.com/redrock/autocrafter/semver"
)

type RecipeType uint8

const (
	BehaviourPack RecipeType = iota + 1
	ResourcePack
	Addon
)

type UUIDPack struct {
	Single string
	RP     string
	BP     string
}

type Recipe struct {
	/* Config */
	Type     RecipeType
	Artifact string

	/* Header */
	Version          *semver.Version
	UUIDs            UUIDPack
	MinEngineVersion *semver.Version

	/* Modules */
	Modules []Module

	/* Meta */
	Authors []string
	License string
}

func (recipe *Recipe) UnmarshalJSON(data []byte) error {
	format_version, json_err := getformatv.Get(data)
	if json_err != nil {
		return json_err
	}

	switch format_version {
	case 1:
		{
			var rawRecipe v1.Recipe
			json_err := json.Unmarshal(data, &rawRecipe)
			if json_err != nil {
				return json_err
			}

			recipe.Type = RecipeType(rawRecipe.Config.Type)
			recipe.Artifact = rawRecipe.Config.Artifact

			recipe.Version = rawRecipe.Header.Version
			recipe.UUIDs.Single = rawRecipe.Header.UUID
			recipe.UUIDs.BP = rawRecipe.Header.UUIDs.BP
			recipe.UUIDs.RP = rawRecipe.Header.UUIDs.RP
			recipe.MinEngineVersion = rawRecipe.Header.MinEngineVersion

			mod_len := len(rawRecipe.Modules)
			recipe.Modules = make([]Module, mod_len)
			for i := 0; i < mod_len; i++ {
				recipe.Modules[i].Description = rawRecipe.Modules[i].Description
				recipe.Modules[i].Type = ModuleType(rawRecipe.Modules[i].Type)
				recipe.Modules[i].Version = rawRecipe.Modules[i].Version
				recipe.Modules[i].UUID = rawRecipe.Modules[i].UUID
			}

			recipe.Authors = rawRecipe.Meta.Authors
			recipe.License = rawRecipe.Meta.License

			return nil
		}
	}

	return &FormatVersionError{msg: fmt.Sprintf("Unknown format version: %d", format_version)}
}

type FormatVersionError struct {
	msg string
}

func (e *FormatVersionError) Error() string {
	return e.msg
}
