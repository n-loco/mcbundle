package recipe

import "github.com/mcbundle/mcbundle/internal/jsonst"

type config struct {
	Type     RecipeType `json:"type"`
	Artifact string     `json:"artifact"`
}

type header struct {
	Name    string          `json:"name"`
	Version *jsonst.SemVer  `json:"version"`
	UUID    *jsonst.UUID    `json:"uuid"`
	UUIDs   [2]*jsonst.UUID `json:"uuids"`
}

type meta struct {
	Authors []string `json:"authors"`
	License string   `json:"license"`
}

type recipeJson struct {
	Config  config   `json:"config"`
	Header  header   `json:"header"`
	Modules []Module `json:"modules"`
	Meta    meta     `json:"meta"`
}
