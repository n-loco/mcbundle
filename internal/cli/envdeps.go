package cli

import (
	"github.com/n-loco/mcbuild/internal/rcontext/commojang"
	"github.com/n-loco/mcbuild/internal/rcontext/recipe"
)

type DependenciesFlags uint8

const (
	ProjectRecipeDependencyFlag DependenciesFlags = 1 << iota
	ComMojangPathDependencyFlag
)

type EnvironmentDependencies struct {
	ProjectRecipe *recipe.Recipe
	ComMojangPath string
}

func GetEnvironmentDependencies(flags DependenciesFlags) EnvironmentDependencies {
	var needsRecipe = (flags & ProjectRecipeDependencyFlag) > 0
	var needsComMojangPath = (flags & ComMojangPathDependencyFlag) > 0
	var deps = EnvironmentDependencies{}

	if needsRecipe {
		deps.ProjectRecipe = recipe.LoadRecipe()
	}

	commojang.WarnComMojangPath(!needsComMojangPath)
	if needsComMojangPath {
		deps.ComMojangPath = commojang.ComMojangPath()
	}

	return deps
}
