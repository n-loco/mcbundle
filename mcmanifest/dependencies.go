package mcmanifest

import (
	"github.com/redrock/autocrafter/jsonst"
	"github.com/redrock/autocrafter/recipe"
)

type Dependency struct {
	UUID       string          `json:"uuid,omitempty,omitzero"`
	ModuleName string          `json:"module_name,omitempty,omitzero"`
	Version    *jsonst.Version `json:"version"`
}

func configAddonDependency(context *MCContext) Dependency {
	projectRecipe := context.Recipe
	category := context.Category

	dependency := Dependency{}
	dependency.Version = projectRecipe.Version

	switch category {
	case recipe.BehavioursCategory:
		dependency.UUID = projectRecipe.UUIDs.RP
	case recipe.ResourcesCategory:
		dependency.UUID = projectRecipe.UUIDs.BP
	}

	return dependency
}
