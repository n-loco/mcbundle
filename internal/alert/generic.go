package alert

import (
	"fmt"
	"reflect"
)

type goErrWrapperAlert struct {
	e error
}

func WrappGoError(err error) Alert {
	var v = reflect.ValueOf(err)
	if err == nil || v.IsNil() {
		return nil
	}
	return &goErrWrapperAlert{e: err}
}

func (errAlert goErrWrapperAlert) Display() string {
	return errAlert.e.Error()
}

func (errAlert goErrWrapperAlert) Tip() string {
	return ""
}

type genericAlert struct {
	msg string
	tip string
}

func AlertF(format string, a ...any) Alert {
	return &genericAlert{msg: fmt.Sprintf(format, a...)}
}

func AlertTF(errFormat string, errA []any, tipFormat string, tipA []any) Alert {
	return &genericAlert{
		msg: fmt.Sprintf(errFormat, errA...),
		tip: fmt.Sprintf(tipFormat, tipA...),
	}
}

func (gAlert genericAlert) Display() string {
	return gAlert.msg
}

func (gAlert genericAlert) Tip() string {
	return gAlert.tip
}
