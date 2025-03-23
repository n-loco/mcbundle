package envdeps

import "github.com/redrock/autocrafter/recipe"

type DependenciesFlags uint8

const (
	ProjectRecipe DependenciesFlags = 1 << iota
	ComMojangPath
)

type EnvironmentDependencies struct {
	ProjectRecipe *recipe.Recipe
	ComMojangPath *string
}

func GetEnvironmentDependencies(flags DependenciesFlags) (EnvironmentDependencies, error) {
	return EnvironmentDependencies{}, nil
}
