package builder

import (
	"encoding/json"
	"os"

	"github.com/redrock/autocrafter/mcmanifest"
	"github.com/redrock/autocrafter/recipe"
)

func GenerateOutput(projectRecipe *recipe.Recipe) {
	switch projectRecipe.Type {
	case recipe.ResourcePackRecipeType:
		{
			resource_mod_dir := os.DirFS("source/contents/resource")
			os.CopyFS("out/", resource_mod_dir)

			manifest := mcmanifest.CreateManifestFromRecipe(projectRecipe, 0)
			data, _ := json.MarshalIndent(manifest, "", "  ")
			os.WriteFile("out/manifest.json", data, os.ModePerm)
		}
	case recipe.BehaviourPackRecipeType:
		{
			data_mod_dir := os.DirFS("source/contents/data")
			os.CopyFS("out/", data_mod_dir)

			manifest := mcmanifest.CreateManifestFromRecipe(projectRecipe, 0)
			data, _ := json.MarshalIndent(manifest, "", "  ")
			os.WriteFile("out/manifest.json", data, os.ModePerm)
		}
	case recipe.AddonRecipeType:
		{
			data_mod_dir := os.DirFS("source/contents/data")
			os.CopyFS("out/bp", data_mod_dir)

			bpManifest := mcmanifest.CreateManifestFromRecipe(projectRecipe, recipe.BehavioursCategory)
			bpData, _ := json.MarshalIndent(bpManifest, "", "  ")
			os.WriteFile("out/bp/manifest.json", bpData, os.ModePerm)

			resource_mod_dir := os.DirFS("source/contents/resource")
			os.CopyFS("out/rp", resource_mod_dir)

			rpManifest := mcmanifest.CreateManifestFromRecipe(projectRecipe, recipe.ResourcesCategory)
			rpData, _ := json.MarshalIndent(rpManifest, "", "  ")
			os.WriteFile("out/rp/manifest.json", rpData, os.ModePerm)
		}
	}
}
