package main

import (
	"github.com/redrock/autocrafter/cli"
	"github.com/redrock/autocrafter/envdeps"
)

func main() {
	cli.SetUpANSIEscapeCodes()
	cli.Print("\x1b[1;92mCOLOR!!\x1b[0m 😊\n")
	cli.Printf("%s\n", envdeps.GetEnvironmentDependencies(envdeps.ComMojangPath).ComMojangPath)
}
