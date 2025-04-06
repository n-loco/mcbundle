package manifest

import (
	"fmt"

	"github.com/n-loco/mcbuild/internal/jsonst"
)

type PackScope byte

const (
	PackScopeAny PackScope = iota
	PackScopeWorld
	PackScopeGlobal
)

func (packScope PackScope) String() string {
	switch packScope {
	case PackScopeAny:
		return "any"
	case PackScopeWorld:
		return "world"
	case PackScopeGlobal:
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
