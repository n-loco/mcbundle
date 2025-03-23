package recipe

import "github.com/redrock/autocrafter/semver"

type ModuleCategory uint8

const (
	Behaviour ModuleCategory = iota + 1
	Resources
)

type ModuleType uint8

const (
	Data ModuleType = iota + 1
	Server
	Resource
)

type Module struct {
	Description string
	Type        ModuleType
	Version     *semver.Version
	UUID        string
}

func (m *Module) Category() ModuleCategory {
	switch m.Type {
	case Data:
	case Server:
		return Behaviour
	case Resource:
		return Resources
	}

	return 0
}
