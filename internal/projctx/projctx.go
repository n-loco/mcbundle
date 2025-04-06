package projctx

import (
	"os"
	"path/filepath"

	"github.com/n-loco/mcbuild/internal/projctx/commojang"
	"github.com/n-loco/mcbuild/internal/projctx/recipe"
)

type EnvRequireFlags byte

const (
	EnvRequireFlagRecipe EnvRequireFlags = 1 << iota
	EnvRequireFlagComMojang
)

type ProjectContext struct {
	Recipe       *recipe.Recipe
	ComMojangDir string
	WorkDir      string
	DistDir      string
	SourceDir    string
}

func CreateProjectContext(flags EnvRequireFlags) (projCtx ProjectContext) {
	var needsRecipe = (flags & EnvRequireFlagRecipe) > 0
	var needsComMojangPath = (flags & EnvRequireFlagComMojang) > 0

	if needsRecipe {
		projCtx.Recipe = recipe.LoadRecipe()

		workDir, wdErr := os.Getwd()

		if wdErr != nil {
			panic("TODO ERRH: " + wdErr.Error())
		}

		projCtx.WorkDir = workDir

		projCtx.DistDir = filepath.Join(workDir, "dist")
		projCtx.SourceDir = filepath.Join(workDir, "source")
	}

	commojang.WarnComMojangPath(!needsComMojangPath)
	if needsComMojangPath {
		projCtx.ComMojangDir = commojang.ComMojangPath()
	}

	return
}
