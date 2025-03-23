package v1

import (
	"encoding/json"
	"fmt"

	"github.com/redrock/autocrafter/semver"
)

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

func (moduleType *ModuleType) UnmarshalJSON(data []byte) error {
	var jstr string

	jsonErr := json.Unmarshal(data, &jstr)

	if jsonErr != nil {
		return jsonErr
	}

	switch jstr {
	case "data":
		*moduleType = Data
		return nil
	case "server":
		*moduleType = Server
		return nil
	case "resource":
		*moduleType = Resource
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
	Description string          `json:"description"`
	Type        ModuleType      `json:"type"`
	Version     *semver.Version `json:"version"`
	UUID        string          `json:"uuid"`
}

func (module *Module) Category() ModuleCategory {
	switch module.Type {
	case Data:
	case Server:
		return Behaviour
	case Resource:
		return Resources
	}

	return 0
}
