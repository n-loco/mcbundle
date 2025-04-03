package rcontext

import "github.com/redrock/autocrafter/rcontext/recipe"

type Context struct {
	Recipe        *recipe.Recipe
	ComMojangPath string
	PackType      recipe.PackType
	Release       bool
}
