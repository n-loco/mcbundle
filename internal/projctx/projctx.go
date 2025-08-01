package projctx

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/assets"
	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

type EnvRequireFlags byte

const (
	EnvRequireFlagRecipe EnvRequireFlags = 1 << iota
	EnvRequireFlagComMojang
)

type ProjectContext struct {
	Diagnostic   alert.Diagnostic
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

func CreateProjectContext(flags EnvRequireFlags, diagnostic alert.Diagnostic) (projCtx ProjectContext) {
	var needsRecipe = (flags & EnvRequireFlagRecipe) > 0
	var needsComMojangPath = (flags & EnvRequireFlagComMojang) > 0

	projCtx.Diagnostic = diagnostic
	projCtx.PackIcon = getPackIcon(diagnostic)

	workDir, getwdErr := os.Getwd()
	if getwdErr != nil {
		diagnostic.AppendError(alert.WrappGoError(getwdErr))
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
			diagnostic.AppendError(alert.WrappGoError(err))
		}
	}

	warnComMojangPath(!needsComMojangPath, diagnostic)
	if needsComMojangPath {
		var dir = comMojangPath(diagnostic)

		projCtx.ComMojangDir = dir
	}

	return
}

func getPackIcon(diagnostic alert.Diagnostic) (iconData []byte) {
	iconData = assets.DefaultPackIcon

	var findPackIconSuccess uint8

	var firstAttemptDiag = getPackIconAttempt("pack_icon.png", &findPackIconSuccess, &iconData)
	var secondAttemptDiag = getPackIconAttempt("icon.png", &findPackIconSuccess, &iconData)

	if findPackIconSuccess > 1 {
		diagnostic.AppendWarning(
			alert.AlertTF(
				"too many icon files, reverting to default", nil,
				"you can use either an icon.png or a pack_icon.png, but not both", nil,
			),
		)
		iconData = assets.DefaultPackIcon
	} else {
		diagnostic.Append(firstAttemptDiag)
		diagnostic.Append(secondAttemptDiag)
	}

	return
}

func getPackIconAttempt(filePath string, findPackIconSuccess *uint8, iconData *[]byte) (diagnostic alert.Diagnostic) {
	diagnostic = alert.NewDiagnostic()

	if packIcon, _ := os.Stat(filePath); packIcon != nil && !packIcon.IsDir() {
		var packIconData, err = os.ReadFile(filePath)
		if err != nil {
			diagnostic.AppendWarning(alert.WrappGoError(err))
		} else {
			*findPackIconSuccess += 1
			*iconData = packIconData
		}
	}
	return
}
