package mcmanifest

import "github.com/redrock/autocrafter/recipe"

type MCManifest struct {
	FormatVersion uint8        `json:"format_version"`
	Header        *Header      `json:"header"`
	Modules       []Module     `json:"modules"`
	Dependencies  []Dependency `json:"dependencies,omitempty,omitzero"`
	Meta          *Meta        `json:"meta,omitempty,omitzero"`
}

func CreateManifestFromRecipe(projectRecipe *recipe.Recipe, filter recipe.Category) *MCManifest {
	var mcManifest *MCManifest = new(MCManifest)
	mcManifest.FormatVersion = 2

	mcManifest.Header = new(Header)
	mcManifest.Header.Description = "pack.description"
	mcManifest.Header.Name = "pack.name"
	mcManifest.Header.Version = projectRecipe.Version
	mcManifest.Header.MinEngineVersion = projectRecipe.MinEngineVersion

	mcManifest.Meta = new(Meta)
	mcManifest.Meta.Authors = projectRecipe.Authors
	mcManifest.Meta.License = projectRecipe.License

	if projectRecipe.Type == recipe.AddonRecipeType {
		mcManifest.Dependencies = make([]Dependency, 1, 6)
		mcManifest.Dependencies[0].Version = projectRecipe.Version

		if filter == recipe.BehavioursCategory {
			mcManifest.Header.UUID = projectRecipe.UUIDs.BP
			mcManifest.Dependencies[0].UUID = projectRecipe.UUIDs.RP
		}
		if filter == recipe.ResourcesCategory {
			mcManifest.Header.UUID = projectRecipe.UUIDs.RP
			mcManifest.Header.PackScope = WorldPackScope
			mcManifest.Dependencies[0].UUID = projectRecipe.UUIDs.BP
		}
	} else {
		mcManifest.Header.UUID = projectRecipe.UUIDs.Single
	}

	return mcManifest
}
