package recipe

import (
	"github.com/redrock/autocrafter/internal/jsonst"
)

type RecipeModule struct {
	Description string
	Type        RecipeModuleType
	Version     *jsonst.SemVer
	UUID        *jsonst.UUID
}
