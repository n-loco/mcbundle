package projfiles

import (
	"fmt"

	"github.com/mcbundle/mcbundle/internal/jsonst"
)

type config struct {
	Type     RecipeType `json:"type"`
	Artifact string     `json:"artifact"`
}

type header struct {
	Name    string          `json:"name"`
	Version *jsonst.SemVer  `json:"version"`
	UUID    *jsonst.UUID    `json:"uuid"`
	UUIDs   [2]*jsonst.UUID `json:"uuids"`
}

type meta struct {
	Authors []string `json:"authors"`
	License string   `json:"license"`
}

type recipeJson struct {
	Config  config   `json:"config"`
	Header  header   `json:"header"`
	Modules []Module `json:"modules"`
	Meta    meta     `json:"meta"`
}

func validateRecipe(rJ *recipeJson) (err error) {
	if rJ.Config.Type == 0 {
		return fmt.Errorf("recipe.json: type is not set")
	}

	if rJ.Config.Artifact == "" {
		return fmt.Errorf("recipe.json: artifact is empty")
	}

	var recipeType = rJ.Config.Type

	if recipeType == RecipeTypeAddOn {
		var uuids = rJ.Header.UUIDs
		if uuids[0] == nil || uuids[1] == nil {
			return fmt.Errorf("recipe.json: uuids are not set correctly")
		}
	} else {
		var uuid = rJ.Header.UUID
		if uuid == nil {
			return fmt.Errorf("recipe.json: uuid is not set")
		}
	}

	if len(rJ.Header.Name) == 0 {
		return fmt.Errorf("recipe.json: name is empty")
	}

	if rJ.Header.Version == nil {
		return fmt.Errorf("recipe.json: version is not set")
	}

	var setModules [3]bool
	for i, mod := range rJ.Modules {
		if mod.Type == 0 {
			return fmt.Errorf("recipe.json: type of module [%d] is not set", i)
		}

		if !recipeType.AcceptsModule(&mod) {
			return fmt.Errorf("recipe.json: module \"%s\" is invalid here", mod.Type.String())
		}

		var modI = uint8((mod.Type) - 1)
		if setModules[modI] {
			return fmt.Errorf("recipe.json: module \"%s\" is set two times", mod.Type.String())
		} else {
			setModules[modI] = true
		}

		if mod.Version == nil {
			return fmt.Errorf("recipe.json: version of module \"%s\" is not set", mod.Type.String())
		}

		if mod.UUID == nil {
			return fmt.Errorf("recipe.json: uuid of module \"%s\" is not set", mod.Type.String())
		}
	}

	return
}
