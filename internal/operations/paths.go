package operations

import (
	"path/filepath"

	"github.com/n-loco/mcbuild/internal/rcontext"
	"github.com/n-loco/mcbuild/internal/rcontext/recipe"
)

const outPath = "dist"

const inPath = "source"

const baseBuildPath = outPath + string(filepath.Separator) + "._obj"

func buildPathRoot(ctx *rcontext.Context) string {
	buildPath := baseBuildPath

	if ctx.Release {
		buildPath = filepath.Join(buildPath, "release")
	} else {
		buildPath = filepath.Join(buildPath, "debug")
	}

	return buildPath
}

func buildPath(ctx *rcontext.Context) string {
	buildPath := buildPathRoot(ctx)

	if ctx.Recipe.Type == recipe.RecipeTypeAddon {
		buildPath = filepath.Join(buildPath, ctx.PackType.Abbr())
	}

	return buildPath
}

func moduleSourcePath(recipeModule *recipe.RecipeModule) string {
	return filepath.Join(inPath, recipeModule.Type.String())
}
