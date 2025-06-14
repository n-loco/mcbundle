package projctx

import (
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/jsonst"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

func packContext(projCtx *ProjectContext, packType projfiles.PackType, release bool) (packCtx PackContext) {
	projRecipe := projCtx.Recipe

	packCtx.ProjectContext = projCtx
	packCtx.PackType = packType
	packCtx.scriptDependencies = make(map[string]*ScriptDependency)
	packCtx.Release = release

	baseDir := filepath.Join(projCtx.DistDir, "build")

	if release {
		baseDir = filepath.Join(baseDir, "release")
	} else {
		baseDir = filepath.Join(baseDir, "debug")
	}

	if projRecipe.Type == projfiles.RecipeTypeAddOn {
		packCtx.PackDistDir = filepath.Join(baseDir, packType.Abbr())
	} else {
		packCtx.PackDistDir = baseDir
	}

	packCtx.PackDirName = projRecipe.DirName()

	if projRecipe.Type == projfiles.RecipeTypeAddOn {
		packCtx.PackDirName += "_" + packType.Abbr()
	}

	if len(projCtx.ComMojangDir) > 0 {
		packCtx.PackDevDir = filepath.Join(packCtx.ComMojangDir, "development_"+packType.ComMojangDirName(), packCtx.PackDirName)
	}

	return
}

func (projCtx *ProjectContext) PackContext(release bool) (packCtx PackContext) {
	packCtx = packContext(projCtx, projCtx.Recipe.PackType(), release)
	return
}

func (projCtx *ProjectContext) AddonContext(release bool) (bpCtx PackContext, rpCtx PackContext) {
	bpCtx = packContext(projCtx, projfiles.PackTypeBehavior, release)
	rpCtx = packContext(projCtx, projfiles.PackTypeResources, release)
	return
}

type ScriptDependency struct {
	Name    string
	Version *jsonst.SemVer
}

type PackContext struct {
	*ProjectContext
	PackType    projfiles.PackType
	Release     bool
	PackDistDir string
	PackDirName string
	PackDevDir  string

	scriptDependencies map[string]*ScriptDependency
}

func (packCtx *PackContext) HasScriptDependency(name string) bool {
	_, has := packCtx.scriptDependencies[name]
	return has
}

func (packCtx *PackContext) AddScriptDependency(name string, version *jsonst.SemVer) {
	has := packCtx.HasScriptDependency(name)

	if !has {
		packCtx.scriptDependencies[name] = &ScriptDependency{Name: name, Version: version}
	}
}

func (packCtx *PackContext) ScriptDependencies() (deps []*ScriptDependency) {
	for _, dep := range packCtx.scriptDependencies {
		deps = append(deps, dep)
	}
	return
}
