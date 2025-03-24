package main

import (
	"github.com/redrock/autocrafter/ansi"
	"github.com/redrock/autocrafter/builder"
	"github.com/redrock/autocrafter/cli"
	"github.com/redrock/autocrafter/envdeps"
)

func main() {
	ansi.SetUpANSIFormatCodes()
	cli.Print("\x1b[1;92mCOLOR!!\x1b[0m 😊\n")
	var deps = envdeps.GetEnvironmentDependencies(envdeps.ProjectRecipe)
	builder.GenerateOutput(deps.ProjectRecipe)
}
