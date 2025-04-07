package v1

import (
	"encoding/json"
	"fmt"
)

type RecipeType uint8

const (
	BehaviorPackRecipeType RecipeType = iota + 1
	ResourcePackRecipeType
	AddonRecipeType
)

func (recipeType RecipeType) String() string {
	switch recipeType {
	case BehaviorPackRecipeType:
		return "Behavior Pack"
	case ResourcePackRecipeType:
		return "Resource Pack"
	case AddonRecipeType:
		return "Addon"
	}
	return ""
}

type Config struct {
	Type     RecipeType
	MojangID string
	Artifact string
}

func (recipeConfig *Config) UnmarshalJSON(data []byte) error {
	var rawRecipeConfig rawConfig

	jsonErr := json.Unmarshal(data, &rawRecipeConfig)

	if jsonErr != nil {
		return jsonErr
	}

	// TODO: missing fields feedback

	var enumErr error

	recipeConfig.Type, enumErr = recipeTypeStringToEnum(rawRecipeConfig.Type)

	if enumErr != nil {
		return enumErr
	}

	recipeConfig.Artifact = rawRecipeConfig.Artifact
	recipeConfig.MojangID = rawRecipeConfig.MojangID

	return nil
}

func (recipeConfig *Config) AllowedCategories() []Category {
	switch recipeConfig.Type {
	case ResourcePackRecipeType:
		return []Category{ResourcesCategory}
	case BehaviorPackRecipeType:
		return []Category{BehavioursCategory}
	case AddonRecipeType:
		return []Category{ResourcesCategory, BehavioursCategory}
	}
	return nil
}

type rawConfig struct {
	Type     string `json:"type"`
	MojangID string `json:"mojang_id"`
	Artifact string `json:"artifact"`
}

func recipeTypeStringToEnum(s string) (RecipeType, error) {
	if len(s) == 0 {
		return 0, &RecipeTypeError{"config.type cannot be empty"}
	}

	switch s {
	case "behavior_pack":
		return BehaviorPackRecipeType, nil
	case "resource_pack":
		return ResourcePackRecipeType, nil
	case "addon":
		return AddonRecipeType, nil
	}

	return 0, &RecipeTypeError{fmt.Sprintf(`In field config.type: unknown type: "%s"`, s)}
}

type RecipeTypeError struct {
	msg string
}

func (err *RecipeTypeError) Error() string {
	return err.msg
}
