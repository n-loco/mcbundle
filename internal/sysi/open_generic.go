//go:build !windows && !linux && !darwin

package sysi

func HasOpenSupport() bool {
	return false
}

func Open(uri string) (err error) {
	panic(unsupportedOpenMessage)
}
