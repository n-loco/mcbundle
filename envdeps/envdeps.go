package envdeps

import "github.com/redrock/autocrafter/recipe"

type DependenciesFlags uint8

const (
	ProjectRecipe DependenciesFlags = 1 << iota
	ComMojangPath
)

type EnvironmentDependencies struct {
	ProjectRecipe *recipe.Recipe
	ComMojangPath string
}

func GetEnvironmentDependencies(flags DependenciesFlags) EnvironmentDependencies {
	var needsRecipe = (flags & ProjectRecipe) > 0
	var needsComMojangPath = (flags & ComMojangPath) > 0

	var deps = EnvironmentDependencies{}

	warnComMojangPath(!needsComMojangPath)
	if needsComMojangPath {
		deps.ComMojangPath = getComMojangPath()
	}

	if needsRecipe {
		deps.ProjectRecipe = recipe.LoadRecipe()
	}

	return deps
}
