package distops

import (
	"path/filepath"

	"github.com/redrock/autocrafter/mcmanifest"
	"github.com/redrock/autocrafter/recipe"
)

type ModuleFSType uint8

const (
	ContentsFSType ModuleFSType = iota + 1
	ScriptingFSType
)

func GetModuleFSType(moduleType recipe.ModuleType) ModuleFSType {
	switch moduleType {
	case recipe.ServerModuleType:
		return ScriptingFSType
	case recipe.DataModuleType:
		fallthrough
	case recipe.ResourceModuleType:
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

func GetModuleSourcePath(rMod *recipe.Module) (ModuleFSType, string) {
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

func GetPackDistPath(context *mcmanifest.MCContext) string {
	projectRecipe := context.Recipe
	category := context.Category
	release := context.Release

	if projectRecipe.Type == recipe.AddonRecipeType {
		var packPath string

		switch category {
		case recipe.ResourcesCategory:
			packPath = "rp"
		case recipe.BehavioursCategory:
			packPath = "bp"
		}

		return filepath.Join(DistTypePath(release), packPath)
	}

	return DistTypePath(release)
}
