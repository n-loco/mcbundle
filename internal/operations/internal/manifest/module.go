package manifest

import (
	"fmt"

	"github.com/mcbundle/mcbundle/internal/jsonst"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

type ModuleType byte

const (
	ModuleTypeData ModuleType = iota + 1
	ModuleTypeScript
	ModuleTypeResources
)

func (moduleType ModuleType) String() string {
	switch moduleType {
	case ModuleTypeResources:
		return "resources"
	case ModuleTypeData:
		return "data"
	case ModuleTypeScript:
		return "script"
	}
	return ""
}

func (moduleType ModuleType) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%v"`, moduleType), nil
}

type Module struct {
	Description string         `json:"description,omitempty,omitzero"`
	Type        ModuleType     `json:"type"`
	UUID        *jsonst.UUID   `json:"uuid"`
	Version     *jsonst.SemVer `json:"version"`
	Language    string         `json:"language,omitempty,omitzero"`
	Entry       string         `json:"entry,omitempty,omitzero"`
}

func ModuleTypeFromRecipeModuleType(recipeModType projfiles.ModuleType) ModuleType {
	switch recipeModType {
	case projfiles.ModuleTypeData:
		return ModuleTypeData
	case projfiles.ModuleTypeResources:
		return ModuleTypeResources
	case projfiles.ModuleTypeServer:
		return ModuleTypeScript
	}
	panic("invalid RecipeModuleType")
}
