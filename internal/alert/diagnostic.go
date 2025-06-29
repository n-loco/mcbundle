package alert

import (
	"reflect"
)

type diagnosticInternal struct {
	Warnings []Alert
	Errors   []Alert
}

type Diagnostic struct {
	*diagnosticInternal
}

func NewDiagnostic() Diagnostic {
	return Diagnostic{new(diagnosticInternal)}
}

func (diagnostic Diagnostic) Append(other Diagnostic) {
	for _, warn := range other.Warnings {
		diagnostic.AppendWarning(warn)
	}

	for _, err := range other.Errors {
		diagnostic.AppendError(err)
	}
}

func (diagnostic Diagnostic) AppendWarning(warning Alert) {
	if warning == nil || reflect.ValueOf(warning).IsNil() {
		return
	}

	diagnostic.Warnings = append(diagnostic.Warnings, warning)
}

func (diagnostic Diagnostic) AppendError(err Alert) {
	if err == nil || reflect.ValueOf(err).IsNil() {
		return
	}

	diagnostic.Errors = append(diagnostic.Errors, err)
}

func (diagnostic Diagnostic) IsZero() bool {
	return (len(diagnostic.Warnings) == 0) && (len(diagnostic.Errors) == 0)
}

func (diagnostic Diagnostic) HasErrors() bool {
	return len(diagnostic.Errors) > 0
}
