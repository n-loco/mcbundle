package v1

import (
	"encoding/json"
	"slices"
)

type Recipe struct {
	Config  *Config
	Header  *Header
	Modules []Module
	Meta    *Meta
}

func (recipe *Recipe) UnmarshalJSON(data []byte) error {
	var rawRecipeData *rawRecipe = &rawRecipe{}
	jsonErr := json.Unmarshal(data, rawRecipeData)

	if jsonErr != nil {
		return jsonErr
	}

	if rawRecipeData.Config == nil {
		return &MissingRequiredFieldError{"config"}
	}

	if rawRecipeData.Header == nil {
		return &MissingRequiredFieldError{"header"}
	}

	if rawRecipeData.Modules == nil {
		return &MissingRequiredFieldError{"modules"}
	}

	allowedCategories := rawRecipeData.Config.AllowedCategories()
	for _, mod := range rawRecipeData.Modules {
		modCategory := mod.Category()
		isNotAllowed := slices.Index(allowedCategories, modCategory) == -1

		if isNotAllowed {
			return &InvalidModuleCategoryError{
				mod.Type,
				modCategory,
				allowedCategories,
				rawRecipeData.Config.Type,
			}
		}
	}

	recipe.Config = rawRecipeData.Config
	recipe.Header = rawRecipeData.Header
	recipe.Modules = rawRecipeData.Modules
	recipe.Meta = rawRecipeData.Meta

	return nil
}

type rawRecipe struct {
	Config  *Config  `json:"config"`
	Header  *Header  `json:"header"`
	Modules []Module `json:"modules"`
	Meta    *Meta    `json:"meta"`
}
