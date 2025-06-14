package projfiles

type PackType byte

const (
	PackTypeBehavior PackType = iota + 1
	PackTypeResources
)

func (packType PackType) ComMojangDirName() string {
	switch packType {
	case PackTypeBehavior:
		return "behavior_packs"
	case PackTypeResources:
		return "resource_packs"
	}
	return ""
}

func (packType PackType) Abbr() string {
	switch packType {
	case PackTypeBehavior:
		return "bp"
	case PackTypeResources:
		return "rp"
	}
	return ""
}

func (packType PackType) String() string {
	switch packType {
	case PackTypeBehavior:
		return "behavior"
	case PackTypeResources:
		return "resources"
	}
	return ""
}
