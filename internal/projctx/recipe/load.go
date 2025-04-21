package recipe

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/alert"
)

type SyntaxError struct {
	Offset  int64
	File    string
	Data    []byte
	prevMsg string
}

func (syntaxError *SyntaxError) Error() string {
	var line int64 = 1
	var col int64 = 0

	for i, c := range syntaxError.Data {
		if i == int(syntaxError.Offset) {
			break
		}

		if c == '\n' {
			col = 0
			line++
		} else {
			col++
		}
	}

	return fmt.Sprintf("%s:%d:%d: %s", syntaxError.File, line, col, syntaxError.prevMsg)
}

func LoadRecipe(workingDir string) (projRecipe *Recipe, diagnostic *alert.Diagnostic) {
	fileData, fileErr := os.ReadFile(filepath.Join(workingDir, "recipe.json"))

	if fileErr != nil {
		diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(fileErr))
		return
	}

	vRecipe := new(Recipe)

	if jsonErr := json.Unmarshal(fileData, vRecipe); jsonErr != nil {
		switch err := jsonErr.(type) {
		case *json.SyntaxError:
			jsonErr = &SyntaxError{
				Offset:  err.Offset,
				File:    "recipe.json",
				Data:    fileData,
				prevMsg: err.Error(),
			}
		}

		diagnostic = diagnostic.AppendError(alert.NewGoErrWrapperAlert(jsonErr))
	}

	projRecipe = vRecipe

	return
}
