package jsonst

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var uuidRegExp = regexp.MustCompile(`([0-9a-fA-F]{8})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{4})-([0-9a-fA-F]{12})`)

type UUID struct {
	bytes [16]byte
}

func (uuid *UUID) String() string {
	blocks := []string{
		bytesToHexStr(uuid.bytes[0:4]),
		bytesToHexStr(uuid.bytes[4:6]),
		bytesToHexStr(uuid.bytes[6:8]),
		bytesToHexStr(uuid.bytes[8:10]),
		bytesToHexStr(uuid.bytes[10:16]),
	}
	return strings.Join(blocks, "-")
}

func (uuid *UUID) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, uuid), nil
}

func (uuid *UUID) UnmarshalJSON(data []byte) error {
	var uuidStr string
	if jsonErr := json.Unmarshal(data, &uuidStr); jsonErr != nil {
		return jsonErr
	}

	if !uuidRegExp.MatchString(uuidStr) {
		return &InvalidUUIDError{String: uuidStr}
	}

	groups := uuidRegExp.FindStringSubmatch(uuidStr)

	strXBytes := make([]string, 0, 16)

	for _, group := range groups {
		for i := 0; i < len(group); i += 2 {
			strXBytes = append(strXBytes, string(rune(group[i]))+string(rune(group[i+1])))
		}
	}

	for i := range 16 {
		b, _ := strconv.ParseUint(strXBytes[i], 16, 8)
		uuid.bytes[i] = byte(b)
	}

	return nil
}

func bytesToHexStr(bytes []byte) string {
	strB := make([]string, 0, len(bytes))
	for _, b := range bytes {
		strB = append(strB, fmt.Sprintf("%02x", b))
	}
	return strings.Join(strB, "")
}
