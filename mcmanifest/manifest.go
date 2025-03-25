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
	mcManifest := new(MCManifest)
	mcManifest.FormatVersion = 2

	mcManifest.Header = createHeaderFromRecipe(projectRecipe, filter)

	mcManifest.Meta = new(Meta)
	mcManifest.Meta.Authors = projectRecipe.Authors
	mcManifest.Meta.License = projectRecipe.License

	if projectRecipe.Type == recipe.AddonRecipeType {
		mcManifest.Dependencies = make([]Dependency, 1, 6)
		mcManifest.Dependencies[0].Version = projectRecipe.Version

		if filter == recipe.BehavioursCategory {
			mcManifest.Dependencies[0].UUID = projectRecipe.UUIDs.RP
		}
		if filter == recipe.ResourcesCategory {
			mcManifest.Dependencies[0].UUID = projectRecipe.UUIDs.BP
		}
	}

	mcManifest.Modules = make([]Module, 0, len(projectRecipe.Modules))
	for _, rMod := range projectRecipe.Modules {
		if rMod.Category() == filter {
			mcManifest.Modules = append(mcManifest.Modules, createModuleFromRecipeModule(rMod))
		}
	}

	return mcManifest
}
