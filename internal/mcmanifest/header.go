package mcmanifest

import (
	"fmt"

	"github.com/n-loco/mcbuild/internal/jsonst"
	"github.com/n-loco/mcbuild/internal/rcontext"
	"github.com/n-loco/mcbuild/internal/rcontext/recipe"
)

type PackScope uint8

const (
	AnyPackScope PackScope = iota
	WorldPackScope
	GlobalPackScope
)

func (packScope PackScope) String() string {
	switch packScope {
	case AnyPackScope:
		return "any"
	case WorldPackScope:
		return "world"
	case GlobalPackScope:
		return "global"
	}

	return ""
}

func (packScope PackScope) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%v"`, packScope), nil
}

type Header struct {
	Description      string         `json:"description,omitempty"`
	MinEngineVersion *jsonst.SemVer `json:"min_engine_version,omitempty"`
	Name             string         `json:"name"`
	PackScope        PackScope      `json:"pack_scope,omitempty"`
	UUID             *jsonst.UUID   `json:"uuid"`
	Version          *jsonst.SemVer `json:"version"`
}

func createHeader(ctx *rcontext.Context) *Header {
	header := new(Header)

	header.Description = "pack.description"
	header.Name = "pack.name"
	header.Version = ctx.Recipe.Version
	header.MinEngineVersion = ctx.Recipe.MinEngineVersion

	if ctx.Recipe.Type == recipe.RecipeTypeAddon {
		if ctx.PackType == recipe.PackTypeBehaviour {
			header.UUID = ctx.Recipe.UUIDs.BP
		}
		if ctx.PackType == recipe.PackTypeResource {
			header.UUID = ctx.Recipe.UUIDs.RP
			header.PackScope = WorldPackScope
		}
	} else {
		header.UUID = ctx.Recipe.UUIDs.Single
	}

	return header
}
