package v1

import (
	"encoding/json"

	"github.com/redrock/autocrafter/jsonst"
)

type UUIDPair struct {
	BP *jsonst.UUID `json:"behaviour_pack"`
	RP *jsonst.UUID `json:"resource_pack"`
}

type Header struct {
	Version          *jsonst.SemVer
	UUID             *jsonst.UUID
	UUIDs            *UUIDPair
	MinEngineVersion *jsonst.SemVer
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
	MinEngineVersion *jsonst.SemVer `json:"min_engine_version"`
}
