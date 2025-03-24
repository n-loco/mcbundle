package mcmanifest

import (
	"fmt"

	"github.com/redrock/autocrafter/semver"
)

type ModuleType uint8

const (
	DataModuleType ModuleType = iota + 1
	ScriptModuleType
	ResourcesModuleType
)

func (moduleType ModuleType) String() string {
	switch moduleType {
	case ResourcesModuleType:
		return "resources"
	case DataModuleType:
		return "data"
	case ScriptModuleType:
		return "script"
	}
	return ""
}

func (moduleType ModuleType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%v"`, moduleType)), nil
}

type Module struct {
	Description string          `json:"description,omitempty,omitzero"`
	Type        ModuleType      `json:"type"`
	UUID        string          `json:"uuid"`
	Version     *semver.Version `json:"version"`
	Language    string          `json:"language,omitempty,omitzero"`
	Entry       string          `json:"entry,omitempty,omitzero"`
}
