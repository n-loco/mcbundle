package jsonst

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type SemVer struct {
	Major      uint8
	Minor      uint8
	Patch      uint8
	Prerelease string
	Build      string
}

var versionRegExp = regexp.MustCompile(`(0|(?:[1-9]\d*))\.(0|(?:[1-9]\d*))\.(0|(?:[1-9]\d*))(?:-((?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*))(?:\.(?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*)))*))?(?:\+((?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*))(?:\.(?:0|(?:[1-9A-Za-z-][0-9A-Za-z-]*)))*))?`)

func (semVer *SemVer) String() string {
	if semVer == nil {
		return ""
	}

	verStr := fmt.Sprintf("%d.%d.%d", semVer.Major, semVer.Minor, semVer.Patch)

	if semVer.Prerelease != "" {
		verStr = fmt.Sprintf("%s-%s", verStr, semVer.Prerelease)
	}

	if semVer.Build != "" {
		verStr = fmt.Sprintf("%s+%s", verStr, semVer.Build)
	}

	return verStr
}

func (v *SemVer) MarshalJSON() ([]byte, error) {
	return fmt.Appendf(nil, `"%s"`, v), nil
}

func (semVer *SemVer) UnmarshalJSON(data []byte) error {
	var vArray [3]uint8
	errArray := json.Unmarshal(data, &vArray)

	if errArray == nil {
		semVer.Major = vArray[0]
		semVer.Minor = vArray[1]
		semVer.Patch = vArray[2]
		return nil
	}

	var vString string
	errString := json.Unmarshal(data, &vString)

	if errString == nil {
		if !versionRegExp.MatchString(vString) {
			return &InvalidSemVerError{String: vString}
		}

		groups := versionRegExp.FindStringSubmatch(vString)

		major, _ := strconv.ParseUint(groups[1], 10, 8)
		minor, _ := strconv.ParseUint(groups[2], 10, 8)
		patch, _ := strconv.ParseUint(groups[3], 10, 8)
		prerelease := groups[4]
		build := groups[5]

		semVer.Major = uint8(major)
		semVer.Minor = uint8(minor)
		semVer.Patch = uint8(patch)

		if prerelease != "" {
			semVer.Prerelease = prerelease
		}

		if build != "" {
			semVer.Build = build
		}

		return nil
	}

	return &json.UnmarshalTypeError{
		Value:  errString.(*json.UnmarshalTypeError).Value,
		Type:   reflect.TypeFor[SemVer](),
		Offset: errString.(*json.UnmarshalTypeError).Offset,
	}
}
