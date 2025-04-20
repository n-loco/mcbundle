package alert

type Alert interface {
	Display() string
	Tip() string
}
