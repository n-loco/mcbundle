package recipe

type RecipeModuleType byte

const (
	RecipeModuleTypeData RecipeModuleType = iota + 1
	RecipeModuleTypeServer
	RecipeModuleTypeResources
)

func (moduleType RecipeModuleType) String() string {
	switch moduleType {
	case RecipeModuleTypeData:
		return "data"
	case RecipeModuleTypeServer:
		return "server"
	case RecipeModuleTypeResources:
		return "resources"
	}
	return ""
}

func (recipeModule RecipeModuleType) PackType() PackType {
	switch recipeModule {
	case RecipeModuleTypeData:
		fallthrough
	case RecipeModuleTypeServer:
		return PackTypeBehaviour
	case RecipeModuleTypeResources:
		return PackTypeResource
	}
	panic("unknown RecipeModuleType value")
}
