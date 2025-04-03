package v1

import (
	"fmt"
	"strings"
)

type MissingRequiredFieldError struct {
	FieldPath string
}

func (err *MissingRequiredFieldError) Error() string {
	return fmt.Sprintf("recipe: missing required field: %s", err.FieldPath)
}

type InvalidModuleCategoryError struct {
	ModuleType        ModuleType
	ModuleCategory    Category
	AllowedCategories []Category
	RecipeType        RecipeType
}

func (err *InvalidModuleCategoryError) Error() string {
	msg := strings.Join([]string{
		fmt.Sprintf("recipe: recipes of type %v only accepts modules of categories: %v;", err.RecipeType, err.AllowedCategories),
		fmt.Sprintf("        but %v modules belongs to the %v category.", err.ModuleType, err.ModuleCategory),
	}, "\n")

	return msg
}
