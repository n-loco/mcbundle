package projctx

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/assets"
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
	PackIcon     []byte
	ComMojangDir string
	WorkDir      string
	DistDir      string
	SourceDir    string
}

func CreateProjectContext(flags EnvRequireFlags) (projCtx ProjectContext, diagnostic *alert.Diagnostic) {
	var needsRecipe = (flags & EnvRequireFlagRecipe) > 0
	var needsComMojangPath = (flags & EnvRequireFlagComMojang) > 0

	projCtx.PackIcon = assets.DefaultPackIcon[:]
	var findPackIconSuccess uint8

	if packIcon, _ := os.Stat("pack_icon.png"); packIcon != nil && !packIcon.IsDir() {
		var packIconData, err = os.ReadFile("pack_icon.png")
		if err != nil {
			diagnostic = diagnostic.AppendWarning(alert.WrappGoError(err))
		} else {
			findPackIconSuccess++
			projCtx.PackIcon = packIconData
		}
	}

	if packIcon, _ := os.Stat("icon.png"); packIcon != nil && !packIcon.IsDir() {
		var packIconData, err = os.ReadFile("icon.png")
		if err != nil {
			diagnostic = diagnostic.AppendWarning(alert.WrappGoError(err))
		} else {
			findPackIconSuccess++
			projCtx.PackIcon = packIconData
		}
	}

	if findPackIconSuccess > 1 {
		diagnostic = diagnostic.AppendWarning(
			alert.AlertTF(
				"too many icon files, reverting to default", nil,
				"you can use either an icon.png or a pack_icon.png, but not both", nil,
			),
		)
		projCtx.PackIcon = assets.DefaultPackIcon[:]
	}

	workDir, getwdErr := os.Getwd()
	if getwdErr != nil {
		diagnostic = diagnostic.AppendError(alert.WrappGoError(getwdErr))
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
			diagnostic = diagnostic.Append(diagnostic.AppendError(alert.WrappGoError(err)))
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
