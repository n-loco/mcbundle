package mcmanifest

import (
	"github.com/redrock/autocrafter/jsonst"
	"github.com/redrock/autocrafter/rcontext"
	"github.com/redrock/autocrafter/rcontext/recipe"
)

type Dependency struct {
	UUID       *jsonst.UUID   `json:"uuid,omitempty,omitzero"`
	ModuleName string         `json:"module_name,omitempty,omitzero"`
	Version    *jsonst.SemVer `json:"version"`
}

func configAddonDependency(context *rcontext.Context) Dependency {
	projectRecipe := context.Recipe
	packType := context.PackType

	dependency := Dependency{}
	dependency.Version = projectRecipe.Version

	switch packType {
	case recipe.PackTypeBehaviour:
		dependency.UUID = projectRecipe.UUIDs.RP
	case recipe.PackTypeResource:
		dependency.UUID = projectRecipe.UUIDs.BP
	}

	return dependency
}
