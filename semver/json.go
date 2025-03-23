package semver

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

const regexStr = `^(0|(?:[1-9]\d*))\.(0|(?:[1-9]\d*))\.(0|(?:[1-9]\d*))(?:-((?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*))(?:\.(?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*)))*))?(?:\+((?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*))(?:\.(?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*)))*))?$`

var vregexp = regexp.MustCompile(regexStr)

func (v *Version) MarshalJSON() ([]byte, error) {
	json_str := fmt.Sprintf(`"%s"`, v)
	return []byte(json_str), nil
}

func (v *Version) UnmarshalJSON(data []byte) error {
	var v_array [3]uint8
	err_array := json.Unmarshal(data, &v_array)

	if err_array == nil {
		v.Major = v_array[0]
		v.Minor = v_array[1]
		v.Patch = v_array[2]

		return nil
	}

	var v_string string
	err_string := json.Unmarshal(data, &v_string)

	if err_string == nil {
		if !vregexp.MatchString(v_string) {
			return &VersionUnmarshalError{msg: "Invalid SemVer."}
		}

		matches := vregexp.FindStringSubmatch(v_string)

		major, _ := strconv.ParseUint(matches[1], 10, 8)
		minor, _ := strconv.ParseUint(matches[2], 10, 8)
		patch, _ := strconv.ParseUint(matches[3], 10, 8)
		prerelease := matches[4]
		build := matches[5]

		v.Major = uint8(major)
		v.Minor = uint8(minor)
		v.Patch = uint8(patch)

		if prerelease != "" {
			v.Prerelease = prerelease
		}

		if build != "" {
			v.Build = build
		}

		return nil
	}

	return &VersionUnmarshalError{msg: `"version" must be a SemVer string or an array of 3 numbers.`}
}
