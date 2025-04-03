package distops

import (
	"path/filepath"

	"github.com/redrock/autocrafter/rcontext"
	"github.com/redrock/autocrafter/rcontext/recipe"
)

type ModuleFSType uint8

const (
	ContentsFSType ModuleFSType = iota + 1
	ScriptingFSType
)

func GetModuleFSType(moduleType recipe.RecipeModuleType) ModuleFSType {
	switch moduleType {
	case recipe.RecipeModuleTypeServer:
		return ScriptingFSType
	case recipe.RecipeModuleTypeData:
		fallthrough
	case recipe.RecipeModuleTypeResources:
		return ContentsFSType
	}
	return 0
}

func (fsType ModuleFSType) String() string {
	switch fsType {
	case ContentsFSType:
		return "contents"
	case ScriptingFSType:
		return "scripting"
	}
	return ""
}

const SourcePath = "source"

func GetModuleSourcePath(rMod *recipe.RecipeModule) (ModuleFSType, string) {
	fsType := GetModuleFSType(rMod.Type)
	superDir := fsType.String()
	subDir := rMod.Type.String()
	return fsType, filepath.Join(SourcePath, superDir, subDir)
}

const DistPath = "dist"

func DistTypePath(release bool) string {
	var buildPath string

	if release {
		buildPath = "release"
	} else {
		buildPath = "debug"
	}

	return filepath.Join("dist", buildPath)
}

func GetPackDistPath(context *rcontext.Context) string {
	projectRecipe := context.Recipe
	category := context.PackType
	release := context.Release

	if projectRecipe.Type == recipe.RecipeTypeAddon {
		var packPath string

		switch category {
		case recipe.PackTypeResource:
			packPath = "rp"
		case recipe.PackTypeBehaviour:
			packPath = "bp"
		}

		return filepath.Join(DistTypePath(release), packPath)
	}

	return DistTypePath(release)
}
