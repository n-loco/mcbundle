package semver

type VersionUnmarshalError struct {
	msg string
}

func (e *VersionUnmarshalError) Error() string {
	return e.msg
}
