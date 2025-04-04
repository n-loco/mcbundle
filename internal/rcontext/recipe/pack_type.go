package recipe

type PackType byte

const (
	PackTypeBehaviour = iota + 1
	PackTypeResource
)

func (packType PackType) ComMojangID() string {
	switch packType {
	case PackTypeBehaviour:
		return "behaviour_pack"
	case PackTypeResource:
		return "resource_pack"
	}
	return ""
}

func (packType PackType) Abbr() string {
	switch packType {
	case PackTypeBehaviour:
		return "bp"
	case PackTypeResource:
		return "rp"
	}
	return ""
}

func (packType PackType) String() string {
	switch packType {
	case PackTypeBehaviour:
		return "behaviour"
	case PackTypeResource:
		return "resource"
	}
	return ""
}
