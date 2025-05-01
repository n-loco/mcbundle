package recipe

type PackType byte

const (
	PackTypeBehavior PackType = iota + 1
	PackTypeResource
)

func (packType PackType) ComMojangID() string {
	switch packType {
	case PackTypeBehavior:
		return "behavior_packs"
	case PackTypeResource:
		return "resource_packs"
	}
	return ""
}

func (packType PackType) Abbr() string {
	switch packType {
	case PackTypeBehavior:
		return "bp"
	case PackTypeResource:
		return "rp"
	}
	return ""
}

func (packType PackType) String() string {
	switch packType {
	case PackTypeBehavior:
		return "behavior"
	case PackTypeResource:
		return "resource"
	}
	return ""
}
