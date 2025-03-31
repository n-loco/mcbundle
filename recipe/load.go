package recipe

import (
	"encoding/json"
	"os"

	"github.com/redrock/autocrafter/terminal"
)

func LoadRecipe() *Recipe {
	data, fileErr := os.ReadFile("recipe.json")
	if fileErr != nil {
		terminal.Eprintf("%s\n", fileErr.Error())
		os.Exit(1)
	}

	recipe := new(Recipe)
	if err := json.Unmarshal(data, recipe); err != nil {
		terminal.Eprintf("%s\n", err.Error())
		os.Exit(1)
	}

	return recipe
}
