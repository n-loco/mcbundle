package rcontext

import "github.com/redrock/autocrafter/internal/rcontext/recipe"

type Context struct {
	Recipe        *recipe.Recipe
	ComMojangPath string
	PackType      recipe.PackType
	Release       bool
}
