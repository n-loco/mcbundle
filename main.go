package main

import (
	"github.com/redrock/autocrafter/cli"
	"github.com/redrock/autocrafter/distops"
	"github.com/redrock/autocrafter/terminal"
)

func main() {
	terminal.SetUpANSIFormatCodes()
	deps := cli.GetEnvironmentDependencies(cli.ProjectRecipeDependencyFlag)
	distops.GeneratePackageTree(deps.ProjectRecipe)
}
