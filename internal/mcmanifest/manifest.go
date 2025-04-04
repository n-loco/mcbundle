package mcmanifest

import (
	"github.com/redrock/autocrafter/internal/rcontext"
)

type MCManifest struct {
	FormatVersion uint8        `json:"format_version"`
	Header        *Header      `json:"header"`
	Modules       []Module     `json:"modules"`
	Dependencies  []Dependency `json:"dependencies,omitempty,omitzero"`
	Meta          *Meta        `json:"meta,omitempty,omitzero"`
}

func CreateManifest(ctx *rcontext.Context) *MCManifest {
	projectRecipe := ctx.Recipe

	mcManifest := new(MCManifest)
	mcManifest.FormatVersion = 2

	mcManifest.Header = createHeader(ctx)

	mcManifest.Meta = new(Meta)
	mcManifest.Meta.Authors = projectRecipe.Authors
	mcManifest.Meta.License = projectRecipe.License

	return mcManifest
}
