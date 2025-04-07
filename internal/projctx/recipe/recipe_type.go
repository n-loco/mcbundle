package recipe

type RecipeType byte

const (
	RecipeTypeBehaviorPack RecipeType = iota + 1
	RecipeTypeResourcePack
	RecipeTypeAddon
)

func (recipeType RecipeType) PackType() PackType {
	switch recipeType {
	case RecipeTypeBehaviorPack:
		return PackTypeBehavior
	case RecipeTypeResourcePack:
		return PackTypeResource
	case RecipeTypeAddon:
		panic("RecipeTypeAddon value cannot have a single PackType")
	}
	panic("unknown RecipeType value")
}
