package mcfiles

import (
	"fmt"

	"github.com/mcbundle/mcbundle/internal/jsonst"
	"github.com/mcbundle/mcbundle/internal/projfiles"
)

type Manifest struct {
	FormatVersion uint8        `json:"format_version"`
	Header        Header       `json:"header"`
	Modules       []Module     `json:"modules"`
	Dependencies  []Dependency `json:"dependencies,omitempty,omitzero"`
	MetaData      MetaData     `json:"metadata,omitempty,omitzero"`
}

// type Manifest

type Header struct {
	Description      string         `json:"description,omitempty"`
	MinEngineVersion [3]uint8       `json:"min_engine_version,omitempty"`
	Name             string         `json:"name"`
	PackScope        PackScope      `json:"pack_scope,omitempty"`
	UUID             *jsonst.UUID   `json:"uuid"`
	Version          *jsonst.SemVer `json:"version"`
}

// type Header

type Module struct {
	Description string         `json:"description,omitempty,omitzero"`
	Type        ModuleType     `json:"type"`
	UUID        *jsonst.UUID   `json:"uuid"`
	Version     *jsonst.SemVer `json:"version"`
	Language    string         `json:"language,omitempty,omitzero"`
	Entry       string         `json:"entry,omitempty,omitzero"`
}

// type Module

type Dependency struct {
	UUID       *jsonst.UUID   `json:"uuid,omitempty,omitzero"`
	ModuleName string         `json:"module_name,omitempty,omitzero"`
	Version    *jsonst.SemVer `json:"version"`
}

// type Dependency

type MetaData struct {
	Authors []string `json:"authors,omitempty,omitzero"`
	License string   `json:"license,omitempty,omitzero"`
}

func (meta *MetaData) IsZero() bool {
	return (len(meta.Authors) == 0) && (len(meta.License) == 0)
}

// type MetaData

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

// type PackScope

type ModuleType byte

const (
	ModuleTypeData ModuleType = iota + 1
	ModuleTypeScript
	ModuleTypeResources
)

func (moduleType ModuleType) String() string {
	switch moduleType {
	case ModuleTypeResources:
		return "resources"
	case ModuleTypeData:
		return "data"
	case ModuleTypeScript:
		return "script"
	}
	return ""
}

func (moduleType ModuleType) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%v"`, moduleType), nil
}

func ModuleTypeFromRecipeModuleType(recipeModType projfiles.ModuleType) ModuleType {
	switch recipeModType {
	case projfiles.ModuleTypeData:
		return ModuleTypeData
	case projfiles.ModuleTypeResources:
		return ModuleTypeResources
	case projfiles.ModuleTypeServer:
		return ModuleTypeScript
	}
	panic("invalid RecipeModuleType")
}

// type ModuleType
