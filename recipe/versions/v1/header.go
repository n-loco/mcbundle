package v1

import (
	"encoding/json"

	"github.com/redrock/autocrafter/semver"
)

type UUIDPair struct {
	BP string `json:"behaviour_pack"`
	RP string `json:"resource_pack"`
}

type Header struct {
	Version          *semver.Version
	UUID             string
	UUIDs            UUIDPair
	MinEngineVersion *semver.Version
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
	Version          *semver.Version `json:"version"`
	UUID             string          `json:"uuid"`
	UUIDs            UUIDPair        `json:"uuids"`
	MinEngineVersion *semver.Version `json:"min_engine_version"`
}
