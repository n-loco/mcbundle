package cli

import "github.com/redrock/autocrafter/distops"

var getTreeTask = TaskDefs{
	Dependencies: ProjectRecipeDependencyFlag,
	Name:         "gen-tree",
	Aliases:      []string{"build"},
	Doc:          "bundles JS/TS and copies content (e. g: data or resources) files into the dist directory.",
	Execute: func(dependencies *EnvironmentDependencies) {
		distops.GeneratePackageTree(dependencies.ProjectRecipe)
	},
}
