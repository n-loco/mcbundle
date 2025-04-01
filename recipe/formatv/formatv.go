package formatv

import (
	"encoding/json"
	"fmt"
)

type formatGetter struct {
	FormatVersion uint8 `json:"format_version"`
}

func Get(data []byte) (uint8, error) {
	var formatg formatGetter

	err := json.Unmarshal(data, &formatg)

	if err != nil {
		return 0, err
	}

	return formatg.FormatVersion, nil
}

type UnsupportedFormatVersionError struct {
	Version uint8
}

func (e UnsupportedFormatVersionError) Error() string {
	return fmt.Sprintf("unsupoted format version: %d", e.Version)
}
