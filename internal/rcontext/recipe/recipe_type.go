package recipe

type RecipeType byte

const (
	RecipeTypeBehaviourPack RecipeType = iota + 1
	RecipeTypeResourcePack
	RecipeTypeAddon
)

func (recipeType RecipeType) PackType() PackType {
	switch recipeType {
	case RecipeTypeBehaviourPack:
		return PackTypeBehaviour
	case RecipeTypeResourcePack:
		return PackTypeResource
	case RecipeTypeAddon:
		panic("RecipeTypeAddon value cannot have a single PackType")
	}
	panic("unknown RecipeType value")
}
