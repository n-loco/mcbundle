package mcmanifest

import "github.com/redrock/autocrafter/semver"

type Dependency struct {
	UUID       string          `json:"uuid,omitempty,omitzero"`
	ModuleName string          `json:"module_name,omitempty,omitzero"`
	Version    *semver.Version `json:"version"`
}
