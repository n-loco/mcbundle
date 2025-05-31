package v1

import (
	"encoding/json"

	"github.com/mcbundle/mcbundle/internal/jsonst"
)

type UUIDPair struct {
	BP *jsonst.UUID `json:"behavior_pack"`
	RP *jsonst.UUID `json:"resource_pack"`
}

type Header struct {
	Version          *jsonst.SemVer
	UUID             *jsonst.UUID
	UUIDs            *UUIDPair
	MinEngineVersion [3]uint8
}

func (recipeHeader *Header) UnmarshalJSON(data []byte) error {
	var rawRecipeHeader rawHeader

	jsonErr := json.Unmarshal(data, &rawRecipeHeader)

	if jsonErr != nil {
		return jsonErr
	}

	// TODO: missing fields feedback

	recipeHeader.Version = rawRecipeHeader.Version
	recipeHeader.UUID = rawRecipeHeader.UUID
	recipeHeader.UUIDs = rawRecipeHeader.UUIDs
	recipeHeader.MinEngineVersion = rawRecipeHeader.MinEngineVersion

	return nil
}

type rawHeader struct {
	Version          *jsonst.SemVer `json:"version"`
	UUID             *jsonst.UUID   `json:"uuid"`
	UUIDs            *UUIDPair      `json:"uuids"`
	MinEngineVersion [3]uint8       `json:"min_engine_version"`
}
