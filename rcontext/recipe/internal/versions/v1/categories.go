package v1

type Category uint8

const (
	BehavioursCategory Category = iota + 1
	ResourcesCategory
)

func (category Category) String() string {
	switch category {
	case BehavioursCategory:
		return "Behaviours"
	case ResourcesCategory:
		return "Resources"
	}
	return ""
}
