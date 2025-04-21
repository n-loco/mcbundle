package alert

type GoErrWrapperAlert struct {
	e error
}

func NewGoErrWrapperAlert(err error) *GoErrWrapperAlert {
	if err == nil {
		return nil
	}
	return &GoErrWrapperAlert{e: err}
}

func (errAlert GoErrWrapperAlert) Display() string {
	return errAlert.e.Error()
}

func (errAlert GoErrWrapperAlert) Tip() string {
	return ""
}
