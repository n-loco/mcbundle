package recipe

import "github.com/redrock/autocrafter/semver"

type Category uint8

const (
	BehavioursCategory Category = iota + 1
	ResourcesCategory
)

type ModuleType uint8

const (
	DataModuleType ModuleType = iota + 1
	ServerModuleType
	ResourceModuleType
)

type Module struct {
	Description string
	Type        ModuleType
	Version     *semver.Version
	UUID        string
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
