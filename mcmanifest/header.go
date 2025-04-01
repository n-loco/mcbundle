package mcmanifest

import (
	"fmt"

	"github.com/redrock/autocrafter/jsonst"
	"github.com/redrock/autocrafter/recipe"
)

type PackScope uint8

const (
	AnyPackScope PackScope = iota
	WorldPackScope
	GlobalPackScope
)

func (packScope PackScope) String() string {
	switch packScope {
	case AnyPackScope:
		return "any"
	case WorldPackScope:
		return "world"
	case GlobalPackScope:
		return "global"
	}

	return ""
}

func (packScope PackScope) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%v"`, packScope), nil
}

type Header struct {
	Description      string          `json:"description,omitempty"`
	MinEngineVersion *jsonst.Version `json:"min_engine_version,omitempty"`
	Name             string          `json:"name"`
	PackScope        PackScope       `json:"pack_scope,omitempty"`
	UUID             string          `json:"uuid"`
	Version          *jsonst.Version `json:"version"`
}

func createHeader(context *MCContext) *Header {
	projectRecipe := context.Recipe
	filter := context.Category

	header := new(Header)

	header.Description = "pack.description"
	header.Name = "pack.name"
	header.Version = projectRecipe.Version
	header.MinEngineVersion = projectRecipe.MinEngineVersion

	if projectRecipe.Type == recipe.AddonRecipeType {
		if filter == recipe.BehavioursCategory {
			header.UUID = projectRecipe.UUIDs.BP
		}
		if filter == recipe.ResourcesCategory {
			header.UUID = projectRecipe.UUIDs.RP
			header.PackScope = WorldPackScope
		}
	} else {
		header.UUID = projectRecipe.UUIDs.Single
	}

	return header
}
