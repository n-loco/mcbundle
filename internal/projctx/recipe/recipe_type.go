package recipe

import (
	"encoding/json"
	"fmt"
)

type RecipeType byte

const (
	RecipeTypeBehaviorPack RecipeType = iota + 1
	RecipeTypeResourcePack
	RecipeTypeAddOn
)

func (recipeType *RecipeType) UnmarshalJSON(data []byte) (err error) {
	var val string
	err = json.Unmarshal(data, &val)
	if err != nil {
		return
	}

	switch val {
	case "behavior_pack":
		*recipeType = RecipeTypeBehaviorPack
	case "resource_pack":
		*recipeType = RecipeTypeResourcePack
	case "addon":
		*recipeType = RecipeTypeAddOn
	default:
		err = fmt.Errorf("unknown recipe type: \"%s\"", val)
	}

	return
}

func (recipeType RecipeType) PackType() PackType {
	switch recipeType {
	case RecipeTypeBehaviorPack:
		return PackTypeBehavior
	case RecipeTypeResourcePack:
		return PackTypeResources
	case RecipeTypeAddOn:
		panic("RecipeTypeAddon value cannot have a single PackType")
	}
	panic("unknown RecipeType value")
}
