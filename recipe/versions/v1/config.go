package v1

import (
	"encoding/json"
	"fmt"
)

type RecipeType uint8

const (
	BehaviourPack RecipeType = iota + 1
	ResourcePack
	Addon
)

type Config struct {
	Type     RecipeType
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

	return nil
}

type rawConfig struct {
	Type     string `json:"type"`
	Artifact string `json:"artifact"`
}

func recipeTypeStringToEnum(s string) (RecipeType, error) {
	if len(s) == 0 {
		return 0, &RecipeTypeError{"config.type cannot be empty"}
	}

	switch s {
	case "behaviour_pack":
		return BehaviourPack, nil
	case "resource_pack":
		return ResourcePack, nil
	case "addon":
		return Addon, nil
	}

	return 0, &RecipeTypeError{fmt.Sprintf(`In field config.type: unknown type: "%s"`, s)}
}

type RecipeTypeError struct {
	msg string
}

func (err *RecipeTypeError) Error() string {
	return err.msg
}
