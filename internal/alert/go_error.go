package alert

import (
	"reflect"
)

type goErrWrapperAlert struct {
	error
}

func WrappGoError(err error) Alert {
	var v = reflect.ValueOf(err)
	if err == nil || v.IsNil() {
		return nil
	}
	return &goErrWrapperAlert{err}
}

func (errAlert goErrWrapperAlert) Display() AlertDisplay {
	return AlertDisplay{
		Message: errAlert.Error(),
	}
}
