package mcmanifest

import (
	"fmt"

	"github.com/redrock/autocrafter/semver"
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
	return []byte(fmt.Sprintf(`"%v"`, packScope)), nil
}

type Header struct {
	Description      string          `json:"description,omitempty"`
	MinEngineVersion *semver.Version `json:"min_engine_version,omitempty"`
	Name             string          `json:"name"`
	PackScope        PackScope       `json:"pack_scope,omitempty"`
	UUID             string          `json:"uuid"`
	Version          *semver.Version `json:"version"`
}
