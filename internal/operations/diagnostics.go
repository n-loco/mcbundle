package operations

import (
	"fmt"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/mcbundle/mcbundle/internal/projctx/recipe"
)

type ESBuildWrapperAlert api.Message

func (alertW *ESBuildWrapperAlert) Display() string {
	file := alertW.Location.File
	line := alertW.Location.Line
	column := alertW.Location.Column
	msg := strings.ToLower(alertW.Text[:1]) + alertW.Text[1:]

	return fmt.Sprintf("%s:%d:%d: %s", file, line, column, msg)
}

func (alertW *ESBuildWrapperAlert) Tip() string {
	return ""
}

type MainFileNotFoundErrAlert struct {
	ExpectedFiles []string
	ModuleType    recipe.RecipeModuleType
}

func (errAlert MainFileNotFoundErrAlert) Display() string {
	if len(errAlert.ExpectedFiles) == 1 {
		return fmt.Sprintf("module %s expected %s as an entry", errAlert.ModuleType, errAlert.ExpectedFiles[0])
	}

	if len(errAlert.ExpectedFiles) == 2 {
		return fmt.Sprintf(
			"module %s expected %s or %s as an entry", errAlert.ModuleType,
			errAlert.ExpectedFiles[0], errAlert.ExpectedFiles[1],
		)
	}

	if len(errAlert.ExpectedFiles) >= 3 {
		return fmt.Sprintf(
			"module %s expected one of these files to be an entry: %s", errAlert.ModuleType,
			strings.Join(errAlert.ExpectedFiles, ", "),
		)
	}

	return ""
}

func (errAlert MainFileNotFoundErrAlert) Tip() string {
	return ""
}
