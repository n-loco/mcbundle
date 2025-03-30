package distops

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/redrock/autocrafter/mcmanifest"
	"github.com/redrock/autocrafter/recipe"
)

func GeneratePackageTree(projectRecipe *recipe.Recipe) {

	if projectRecipe.Type == recipe.AddonRecipeType {
		bpContext := mcmanifest.MCContext{
			Recipe:   projectRecipe,
			Category: recipe.BehavioursCategory,
		}

		genMCPackTree(&bpContext)

		rpContext := mcmanifest.MCContext{
			Recipe:   projectRecipe,
			Category: recipe.ResourcesCategory,
		}

		genMCPackTree(&rpContext)
	} else {
		genMCPackTree(&mcmanifest.MCContext{
			Recipe:   projectRecipe,
			Category: recipe.Any,
		})
	}
}

func genMCPackTree(context *mcmanifest.MCContext) {
	projectRecipe := context.Recipe
	filter := context.Category

	packDistPath := GetPackDistPath(context)

	for _, rMod := range projectRecipe.Modules {
		if filter == recipe.Any || rMod.Category() == filter {
			fsType, sourcePath := GetModuleSourcePath(&rMod)

			if fsType == ContentsFSType {
				modSrcDirFS := os.DirFS(sourcePath)
				os.CopyFS(packDistPath, modSrcDirFS)
			}
		}
	}

	manifest := mcmanifest.CreateManifest(context)
	data, _ := json.MarshalIndent(manifest, "", "  ")
	os.WriteFile(filepath.Join(packDistPath, "manifest.json"), data, os.ModePerm)
}
