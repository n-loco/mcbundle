package cli

import "github.com/n-loco/mcbuild/internal/operations"

var getTreeTask = TaskDefs{
	Dependencies: ProjectRecipeDependencyFlag,
	Name:         "build",
	Doc:          "bundles JS/TS and copies content (e. g: data or resources) files into the dist directory.",
	Execute: func(dependencies *EnvironmentDependencies) {
		operations.BuildProject(dependencies.ProjectRecipe, false)
	},
}
