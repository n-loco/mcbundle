package mcmanifest

import (
	"fmt"

	"github.com/redrock/autocrafter/jsonst"
	"github.com/redrock/autocrafter/rcontext/recipe"
)

type ModuleType uint8

const (
	DataModuleType ModuleType = iota + 1
	ScriptModuleType
	ResourcesModuleType
)

func (moduleType ModuleType) String() string {
	switch moduleType {
	case ResourcesModuleType:
		return "resources"
	case DataModuleType:
		return "data"
	case ScriptModuleType:
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

func createModuleFromRecipeModule(recipeMod recipe.RecipeModule) Module {
	var mod Module

	mod.Description = recipeMod.Description
	mod.UUID = recipeMod.UUID
	mod.Version = recipeMod.Version
	mod.Type = ModuleType(recipeMod.Type)

	if recipeMod.Type == recipe.RecipeModuleTypeServer {
		mod.Language = "javascript"
		mod.Entry = "scripts/server.js"
	}

	return mod
}
