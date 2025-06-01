package recipe

import (
	"encoding/json"
	"fmt"
)

type ModuleType byte

const (
	ModuleTypeData ModuleType = iota + 1
	ModuleTypeServer
	ModuleTypeResources
)

func (moduleType *ModuleType) UnmarshalJSON(data []byte) (err error) {
	var val string
	err = json.Unmarshal(data, &val)
	if err != nil {
		return
	}

	switch val {
	case "data":
		*moduleType = ModuleTypeData
	case "script":
		fallthrough
	case "server":
		*moduleType = ModuleTypeServer
	case "resources":
		*moduleType = ModuleTypeResources
	default:
		err = fmt.Errorf("unknown module type: \"%s\"", val)
	}

	return
}

func (moduleType ModuleType) String() string {
	switch moduleType {
	case ModuleTypeData:
		return "data"
	case ModuleTypeServer:
		return "server"
	case ModuleTypeResources:
		return "resources"
	}
	return ""
}

func (recipeModule ModuleType) BelongsTo() PackType {
	switch recipeModule {
	case ModuleTypeData:
		fallthrough
	case ModuleTypeServer:
		return PackTypeBehavior
	case ModuleTypeResources:
		return PackTypeResources
	}
	panic("unknown ModuleType value")
}
