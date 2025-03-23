package getformatv

import (
	"encoding/json"
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
