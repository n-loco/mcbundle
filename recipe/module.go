package recipe

import (
	"github.com/redrock/autocrafter/jsonst"
)

type Category uint8

const (
	BehavioursCategory Category = iota + 1
	ResourcesCategory
	Any Category = 0xFF
)

type ModuleType uint8

const (
	DataModuleType ModuleType = iota + 1
	ServerModuleType
	ResourceModuleType
)

func (moduleType ModuleType) String() string {
	switch moduleType {
	case DataModuleType:
		return "data"
	case ServerModuleType:
		return "server"
	case ResourceModuleType:
		return "resource"
	}
	return ""
}

type Module struct {
	Description string
	Type        ModuleType
	Version     *jsonst.SemVer
	UUID        *jsonst.UUID
}

func (m *Module) Category() Category {
	switch m.Type {
	case DataModuleType:
		fallthrough
	case ServerModuleType:
		return BehavioursCategory
	case ResourceModuleType:
		return ResourcesCategory
	}

	return 0
}
