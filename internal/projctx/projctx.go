package projctx

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/projctx/commojang"
	"github.com/mcbundle/mcbundle/internal/projctx/recipe"
)

type EnvRequireFlags byte

const (
	EnvRequireFlagRecipe EnvRequireFlags = 1 << iota
	EnvRequireFlagComMojang
)

type ProjectContext struct {
	Recipe       recipe.Recipe
	ComMojangDir string
	WorkDir      string
	DistDir      string
	SourceDir    string
}

func CreateProjectContext(flags EnvRequireFlags) (projCtx ProjectContext, diagnostic *alert.Diagnostic) {
	var needsRecipe = (flags & EnvRequireFlagRecipe) > 0
	var needsComMojangPath = (flags & EnvRequireFlagComMojang) > 0

	workDir, getwdErr := os.Getwd()
	if getwdErr != nil {
		diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(getwdErr))
		return
	}

	projCtx.WorkDir = workDir
	projCtx.DistDir = filepath.Join(workDir, "dist")
	projCtx.SourceDir = filepath.Join(workDir, "source")

	if needsRecipe {
		var recipeFile, err = os.Open("recipe.json")
		defer recipeFile.Close()

		if err == nil {
			err = projCtx.Recipe.Load(recipeFile)
		}

		if err != nil {
			diagnostic = diagnostic.Append(diagnostic.AppendError(alert.NewGoErrWrapperAlert(err)))
		}
	}

	diagnostic = diagnostic.Append(commojang.WarnComMojangPath(!needsComMojangPath))
	if needsComMojangPath {
		dir, comMojangDiag := commojang.ComMojangPath()

		projCtx.ComMojangDir = dir
		diagnostic = diagnostic.Append(comMojangDiag)
	}

	return
}
