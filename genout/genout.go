package genout

import (
	"os"

	"github.com/redrock/autocrafter/recipe"
)

func GenerateOutput(projectRecipe *recipe.Recipe) {
	switch projectRecipe.Type {
	case recipe.ResourcePack:
		{
			resource_mod_dir := os.DirFS("source/contents/resource")
			os.CopyFS("out/rp", resource_mod_dir)
		}
	case recipe.BehaviourPack:
		{
			data_mod_dir := os.DirFS("source/contents/data")
			os.CopyFS("out/bp", data_mod_dir)
		}
	case recipe.Addon:
		{

		}
	}
}
