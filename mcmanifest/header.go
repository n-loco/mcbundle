package mcmanifest

import (
	"fmt"

	"github.com/redrock/autocrafter/recipe"
	"github.com/redrock/autocrafter/semver"
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
	return []byte(fmt.Sprintf(`"%v"`, packScope)), nil
}

type Header struct {
	Description      string          `json:"description,omitempty"`
	MinEngineVersion *semver.Version `json:"min_engine_version,omitempty"`
	Name             string          `json:"name"`
	PackScope        PackScope       `json:"pack_scope,omitempty"`
	UUID             string          `json:"uuid"`
	Version          *semver.Version `json:"version"`
}

func createHeaderFromRecipe(projectRecipe *recipe.Recipe, filter recipe.Category) *Header {
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
