package manifest

import (
	"github.com/mcbundle/mcbundle/internal/jsonst"
)

type Dependency struct {
	UUID       *jsonst.UUID   `json:"uuid,omitempty,omitzero"`
	ModuleName string         `json:"module_name,omitempty,omitzero"`
	Version    *jsonst.SemVer `json:"version"`
}
