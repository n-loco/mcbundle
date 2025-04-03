package distops

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/redrock/autocrafter/mcmanifest"
	"github.com/redrock/autocrafter/rcontext"
	"github.com/redrock/autocrafter/rcontext/recipe"
)

func GeneratePackageTree(projectRecipe *recipe.Recipe) {

	if projectRecipe.Type == recipe.RecipeTypeAddon {
		bpContext := rcontext.Context{
			Recipe:   projectRecipe,
			PackType: recipe.PackTypeBehaviour,
		}

		genMCPackTree(&bpContext)

		rpContext := rcontext.Context{
			Recipe:   projectRecipe,
			PackType: recipe.PackTypeBehaviour,
		}

		genMCPackTree(&rpContext)
	} else {
		genMCPackTree(&rcontext.Context{
			Recipe:   projectRecipe,
			PackType: projectRecipe.Type.PackType(),
		})
	}
}

func genMCPackTree(context *rcontext.Context) {
	projectRecipe := context.Recipe
	filter := context.PackType

	packDistPath := GetPackDistPath(context)

	for _, rMod := range projectRecipe.Modules {
		if rMod.Type.PackType() == filter {
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
