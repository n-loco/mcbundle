package recipe

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/n-loco/bpbuild/internal/terminal"
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

func LoadRecipe() *Recipe {
	data, fileErr := os.ReadFile("recipe.json")
	if fileErr != nil {
		terminal.Eprintf("%s\n", fileErr.Error())
		os.Exit(1)
	}

	recipe := new(Recipe)
	if jsonErr := json.Unmarshal(data, recipe); jsonErr != nil {
		switch err := jsonErr.(type) {
		case *json.SyntaxError:
			jsonErr = &SyntaxError{
				Offset:  err.Offset,
				File:    "recipe.json",
				Data:    data,
				prevMsg: err.Error(),
			}
		}

		terminal.Eprintf("%s\n", jsonErr.Error())
		os.Exit(1)
	}

	return recipe
}
