package operations

import (
	"fmt"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/n-loco/mcbuild/internal/rcontext/recipe"
)

type NativeModuleError struct {
	NativeModule string
}

func (nativeModErr NativeModuleError) Error() string {
	return fmt.Sprintf("unnable to find native module %s's version", nativeModErr.NativeModule)
}

type ESBuildErrorWrapper struct {
	Messages []api.Message
}

func (esberr ESBuildErrorWrapper) Error() string {
	return fmt.Sprintf("%v", esberr.Messages)
}

type MainFileNotFoundError struct {
	ExpectedFiles []string
	ModuleType    recipe.RecipeModuleType
}

func (mfnf MainFileNotFoundError) Error() string {
	if len(mfnf.ExpectedFiles) == 1 {
		return fmt.Sprintf("module %s expected %s as an entry", mfnf.ModuleType, mfnf.ExpectedFiles[0])
	}

	if len(mfnf.ExpectedFiles) == 2 {
		return fmt.Sprintf(
			"module %s expected %s or %s as an entry", mfnf.ModuleType,
			mfnf.ExpectedFiles[0], mfnf.ExpectedFiles[1],
		)
	}

	if len(mfnf.ExpectedFiles) >= 3 {
		return fmt.Sprintf(
			"module %s expected one of these files to be an entry: %s", mfnf.ModuleType,
			strings.Join(mfnf.ExpectedFiles, ", "),
		)
	}

	return ""
}
