package main

import (
	"fmt"

	"github.com/redrock/autocrafter/cli"
	"github.com/redrock/autocrafter/genout"
	"github.com/redrock/autocrafter/recipe"
)

func main() {
	cli.SetUpANSI()
	projectRecipe := recipe.LoadRecipe()
	fmt.Println(projectRecipe)

	genout.GenerateOutput(projectRecipe)
}
