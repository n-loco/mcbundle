package recipe

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func LoadRecipe() *Recipe {
	data, file_err := os.ReadFile("recipe.json")
	if file_err != nil {
		fmt.Fprintln(os.Stderr, file_err)
		fmt.Fprintln(os.Stderr, reflect.TypeOf(file_err))
		os.Exit(1)
	}

	recipe := new(Recipe)
	if err := json.Unmarshal(data, recipe); err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, reflect.TypeOf(err))
		os.Exit(1)
	}

	return recipe
}
