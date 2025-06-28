package projfiles

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/mcbundle/mcbundle/internal/jsonst"
)

type Recipe struct {
	Type     RecipeType
	Artifact string

	Name    string
	Version *jsonst.SemVer
	UUID    *jsonst.UUID
	UUIDs   [2]*jsonst.UUID

	Modules []Module

	Authors []string
	License string
}

func (recipe *Recipe) DirName() string {
	return strings.ReplaceAll(recipe.Name, " ", "")
}

func (recipe *Recipe) Load(reader io.Reader) (err error) {
	var decoder = json.NewDecoder(reader)
	var recipeContent recipeJson
	err = decoder.Decode(&recipeContent)
	if err != nil {
		return
	}
	err = validateRecipe(&recipeContent)
	if err != nil {
		return
	}

	recipe.Type = recipeContent.Config.Type

	recipe.Name = recipeContent.Header.Name
	recipe.Version = recipeContent.Header.Version
	recipe.UUID = recipeContent.Header.UUID
	recipe.UUIDs = recipeContent.Header.UUIDs

	recipe.Modules = recipeContent.Modules

	recipe.Authors = recipeContent.Meta.Authors
	recipe.License = recipeContent.Meta.License

	var vars = variables{make(map[string]string)}

	vars.set("name", ssCodeRegExp.ReplaceAllString(recipeContent.Header.Name, ""))
	vars.set("raw_name", recipeContent.Header.Name)
	vars.set("dir_name", recipe.DirName())
	vars.set("short_name",
		strings.ToLower(strings.ReplaceAll(
			ssCodeRegExp.ReplaceAllString(recipeContent.Header.Name, ""),
			" ", "-",
		)),
	)
	vars.set("version", recipeContent.Header.Version.String())

	recipe.Artifact = vars.apply(recipeContent.Config.Artifact)

	return
}

func (recipe *Recipe) PackType() PackType {
	return recipe.Type.PackType()
}

// type Recipe

type Module struct {
	Type    ModuleType     `json:"type"`
	Version *jsonst.SemVer `json:"version"`
	UUID    *jsonst.UUID   `json:"uuid"`
}

func (module *Module) BelongsTo() PackType {
	return module.Type.BelongsTo()
}

// type Module

type ModuleType byte

const (
	ModuleTypeData ModuleType = iota + 1
	ModuleTypeServer
	ModuleTypeResources
)

func (moduleType *ModuleType) UnmarshalJSON(data []byte) (err error) {
	var val string
	err = json.Unmarshal(data, &val)
	if err != nil {
		return
	}

	switch val {
	case "data":
		*moduleType = ModuleTypeData
	case "script":
		fallthrough
	case "server":
		*moduleType = ModuleTypeServer
	case "resources":
		*moduleType = ModuleTypeResources
	default:
		err = fmt.Errorf("unknown module type: \"%s\"", val)
	}

	return
}

func (moduleType ModuleType) String() string {
	switch moduleType {
	case ModuleTypeData:
		return "data"
	case ModuleTypeServer:
		return "server"
	case ModuleTypeResources:
		return "resources"
	}
	return ""
}

func (recipeModule ModuleType) BelongsTo() PackType {
	switch recipeModule {
	case ModuleTypeData:
		fallthrough
	case ModuleTypeServer:
		return PackTypeBehavior
	case ModuleTypeResources:
		return PackTypeResources
	}
	panic("unknown ModuleType value")
}

// type ModuleType

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

func (recipeType RecipeType) AcceptsModuleType(mT ModuleType) bool {
	if recipeType == RecipeTypeAddOn {
		return ((PackTypeBehavior | PackTypeResources) & mT.BelongsTo()) > 0
	} else {
		return recipeType.PackType() == mT.BelongsTo()
	}
}

func (recipeType RecipeType) AcceptsModule(m *Module) bool {
	return recipeType.AcceptsModuleType(m.Type)
}

// type RecipeType
