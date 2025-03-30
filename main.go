package main

import (
	"github.com/redrock/autocrafter/ansi"
	"github.com/redrock/autocrafter/distops"
	"github.com/redrock/autocrafter/envdeps"
)

func main() {
	ansi.SetUpANSIFormatCodes()
	deps := envdeps.GetEnvironmentDependencies(envdeps.ProjectRecipeDependencyFlag)
	distops.GeneratePackageTree(deps.ProjectRecipe)
}
