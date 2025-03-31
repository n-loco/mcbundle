package main

import (
	"github.com/redrock/autocrafter/distops"
	"github.com/redrock/autocrafter/envdeps"
	"github.com/redrock/autocrafter/terminal"
)

func main() {
	terminal.SetUpANSIFormatCodes()
	deps := envdeps.GetEnvironmentDependencies(envdeps.ProjectRecipeDependencyFlag)
	distops.GeneratePackageTree(deps.ProjectRecipe)
}
