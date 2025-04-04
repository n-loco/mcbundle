package rcontext

import "github.com/n-loco/mcbuild/internal/rcontext/recipe"

type Context struct {
	Recipe        *recipe.Recipe
	ComMojangPath string
	PackType      recipe.PackType
	Release       bool
}
