package jsonst

import "fmt"

type InvalidSemVerError struct {
	String string
}

func (err InvalidSemVerError) Error() string {
	return fmt.Sprintf("invalid semver: %s", err.String)
}

type InvalidUUIDError struct {
	String string
}

func (err InvalidUUIDError) Error() string {
	return fmt.Sprintf("invalid uuid: %s", err.String)
}
