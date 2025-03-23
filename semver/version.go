package semver

import (
	"fmt"
)

type Version struct {
	Major      uint8
	Minor      uint8
	Patch      uint8
	Prerelease string
	Build      string
}

func (v *Version) String() string {
	if v == nil {
		return ""
	}

	verstr := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)

	if v.Prerelease != "" {
		verstr = fmt.Sprintf("%s-%s", verstr, v.Prerelease)
	}

	if v.Build != "" {
		verstr = fmt.Sprintf("%s+%s", verstr, v.Build)
	}

	return verstr
}
