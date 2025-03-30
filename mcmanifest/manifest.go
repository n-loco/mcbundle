package mcmanifest

import "github.com/redrock/autocrafter/recipe"

type MCContext struct {
	Recipe        *recipe.Recipe
	ComMojangPath string
	Category      recipe.Category
	Release       bool
}

type MCManifest struct {
	FormatVersion uint8        `json:"format_version"`
	Header        *Header      `json:"header"`
	Modules       []Module     `json:"modules"`
	Dependencies  []Dependency `json:"dependencies,omitempty,omitzero"`
	Meta          *Meta        `json:"meta,omitempty,omitzero"`
}

func CreateManifest(context *MCContext) *MCManifest {
	projectRecipe := context.Recipe
	filter := context.Category

	mcManifest := new(MCManifest)
	mcManifest.FormatVersion = 2

	mcManifest.Header = createHeader(context)

	mcManifest.Meta = new(Meta)
	mcManifest.Meta.Authors = projectRecipe.Authors
	mcManifest.Meta.License = projectRecipe.License

	if projectRecipe.Type == recipe.AddonRecipeType {
		mcManifest.Dependencies = append(mcManifest.Dependencies, configAddonDependency(context))
	}

	for _, rMod := range projectRecipe.Modules {
		if filter == recipe.Any || rMod.Category() == filter {
			mcManifest.Modules = append(mcManifest.Modules, createModuleFromRecipeModule(rMod))
		}
	}

	return mcManifest
}
