package recipe

import (
	"github.com/n-loco/mcbuild/internal/jsonst"
)

type RecipeModule struct {
	Description string
	Type        RecipeModuleType
	Version     *jsonst.SemVer
	UUID        *jsonst.UUID
}
