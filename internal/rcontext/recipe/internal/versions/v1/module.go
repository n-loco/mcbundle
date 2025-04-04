package v1

import (
	"encoding/json"
	"fmt"

	"github.com/redrock/autocrafter/internal/jsonst"
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
		return "Data"
	case ServerModuleType:
		return "Server"
	case ResourceModuleType:
		return "Resource"
	}
	return ""
}

func (moduleType *ModuleType) UnmarshalJSON(data []byte) error {
	var jstr string

	jsonErr := json.Unmarshal(data, &jstr)

	if jsonErr != nil {
		return jsonErr
	}

	switch jstr {
	case "data":
		*moduleType = DataModuleType
		return nil
	case "server":
		*moduleType = ServerModuleType
		return nil
	case "resource":
		*moduleType = ResourceModuleType
		return nil
	}

	return &ModuleTypeError{fmt.Sprintf(`Unknown type: "%s"`, jstr)}
}

type ModuleTypeError struct {
	msg string
}

func (moduleTypeError *ModuleTypeError) Error() string {
	return moduleTypeError.msg
}

type Module struct {
	Description string         `json:"description"`
	Type        ModuleType     `json:"type"`
	Version     *jsonst.SemVer `json:"version"`
	UUID        *jsonst.UUID   `json:"uuid"`
}

func (module *Module) Category() Category {
	switch module.Type {
	case DataModuleType:
		fallthrough
	case ServerModuleType:
		return BehavioursCategory
	case ResourceModuleType:
		return ResourcesCategory
	}

	return 0
}
