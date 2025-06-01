package recipe

import (
	"encoding/json"
	"io"
	"regexp"
	"strings"

	"github.com/mcbundle/mcbundle/internal/jsonst"
)

var templateRegExp = regexp.MustCompile(`{{([_a-zA-Z][_a-zA-Z0-9]*)}}`)

var ssCodeRegExp = regexp.MustCompile(`§.`)

type variables struct {
	intern map[string]string
}

func (vars *variables) get(varName string) string {
	value, ok := vars.intern[varName]

	if ok {
		return value
	}

	return "null"
}

func (vars *variables) set(varName string, value string) {
	vars.intern[varName] = value
}

type Recipe struct {
	Type     RecipeType
	Artifact string

	Name    string
	Version *jsonst.SemVer
	UUID    *jsonst.UUID
	UUIDs   [2]*jsonst.UUID

	Modules []Module

	Authors []string
	License string
}

func (recipe *Recipe) DirName() string {
	return strings.ReplaceAll(recipe.Name, " ", "")
}

func (recipe *Recipe) Load(reader io.Reader) (err error) {
	var decoder = json.NewDecoder(reader)
	var recipeContent recipeJson
	err = decoder.Decode(&recipeContent)
	if err != nil {
		return
	}

	var vars = variables{intern: make(map[string]string)}

	vars.set("name", ssCodeRegExp.ReplaceAllString(recipeContent.Header.Name, ""))
	vars.set("raw_name", recipeContent.Header.Name)
	vars.set("dir_name", strings.ReplaceAll(recipeContent.Header.Name, " ", ""))
	vars.set("short_name",
		strings.ToLower(strings.ReplaceAll(
			ssCodeRegExp.ReplaceAllString(recipeContent.Header.Name, ""),
			" ", "-",
		)),
	)
	vars.set("version", recipeContent.Header.Version.String())

	recipe.Type = recipeContent.Config.Type
	recipe.Artifact = templateRegExp.ReplaceAllStringFunc(
		recipeContent.Config.Artifact,
		func(value string) string {
			return vars.get(value[2 : len(value)-2])
		})

	recipe.Name = recipeContent.Header.Name
	recipe.Version = recipeContent.Header.Version
	recipe.UUID = recipeContent.Header.UUID
	recipe.UUIDs = recipeContent.Header.UUIDs

	recipe.Modules = recipeContent.Modules

	recipe.Authors = recipeContent.Meta.Authors
	recipe.License = recipeContent.Meta.License

	return
}

func (recipe *Recipe) PackType() PackType {
	return recipe.Type.PackType()
}

type Module struct {
	Type    ModuleType     `json:"type"`
	Version *jsonst.SemVer `json:"version"`
	UUID    *jsonst.UUID   `json:"uuid"`
}

func (module *Module) BelongsTo() PackType {
	return module.Type.BelongsTo()
}
