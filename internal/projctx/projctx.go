package projctx

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/assets"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

type EnvRequireFlags byte

const (
	EnvRequireFlagRecipe EnvRequireFlags = 1 << iota
	EnvRequireFlagComMojang
)

type ProjectContext struct {
	Recipe       projfiles.Recipe
	PackIcon     []byte
	ComMojangDir string
	WorkDir      string
	DistDir      string
	SourceDir    string
}

func (projCtx *ProjectContext) Rel(path string) (string, error) {
	return filepath.Rel(projCtx.WorkDir, path)
}

func CreateProjectContext(flags EnvRequireFlags) (projCtx ProjectContext, diagnostic *alert.Diagnostic) {
	var needsRecipe = (flags & EnvRequireFlagRecipe) > 0
	var needsComMojangPath = (flags & EnvRequireFlagComMojang) > 0

	projCtx.PackIcon, diagnostic = getPackIcon()

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

		if err == nil {
			err = projCtx.Recipe.Load(recipeFile)
			recipeFile.Close()
		}

		if err != nil {
			diagnostic = diagnostic.Append(diagnostic.AppendError(alert.WrappGoError(err)))
		}
	}

	diagnostic = diagnostic.Append(warnComMojangPath(!needsComMojangPath))
	if needsComMojangPath {
		dir, comMojangDiag := comMojangPath()

		projCtx.ComMojangDir = dir
		diagnostic = diagnostic.Append(comMojangDiag)
	}

	return
}

func getPackIcon() (iconData []byte, diagnostic *alert.Diagnostic) {
	iconData = assets.DefaultPackIcon[:]

	var findPackIconSuccess uint8

	var firstAttemptDiag = getPackIconAttempt("pack_icon.png", &findPackIconSuccess, &iconData)
	var secondAttemptDiag = getPackIconAttempt("icon.png", &findPackIconSuccess, &iconData)

	if findPackIconSuccess > 1 {
		diagnostic = diagnostic.AppendWarning(
			alert.AlertTF(
				"too many icon files, reverting to default", nil,
				"you can use either an icon.png or a pack_icon.png, but not both", nil,
			),
		)
		iconData = assets.DefaultPackIcon[:]
	} else {
		diagnostic = diagnostic.Append(firstAttemptDiag)
		diagnostic = diagnostic.Append(secondAttemptDiag)
	}

	return
}

func getPackIconAttempt(filePath string, findPackIconSuccess *uint8, iconData *[]byte) (diagnostic *alert.Diagnostic) {
	if packIcon, _ := os.Stat(filePath); packIcon != nil && !packIcon.IsDir() {
		var packIconData, err = os.ReadFile(filePath)
		if err != nil {
			diagnostic = diagnostic.AppendWarning(alert.WrappGoError(err))
		} else {
			*findPackIconSuccess += 1
			*iconData = packIconData
		}
	}
	return
}
