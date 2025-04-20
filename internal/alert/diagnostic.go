package alert

type Diagnostic struct {
	Warnings []Alert
	Errors   []Alert
}

func (diagnostic *Diagnostic) Append(other *Diagnostic) *Diagnostic {
	if other == nil {
		return diagnostic
	}

	if diagnostic == nil {
		diagnostic = new(Diagnostic)
	}

	diagnostic.Warnings = append(diagnostic.Warnings, other.Warnings...)
	diagnostic.Errors = append(diagnostic.Errors, other.Errors...)

	return diagnostic
}

func (diagnostic *Diagnostic) AppendWarning(warning Alert) *Diagnostic {
	if warning == nil {
		return diagnostic
	}

	if diagnostic == nil {
		diagnostic = new(Diagnostic)
	}

	diagnostic.Warnings = append(diagnostic.Warnings, warning)

	return diagnostic
}

func (diagnostic *Diagnostic) AppendError(err Alert) *Diagnostic {
	if err == nil {
		return diagnostic
	}

	if diagnostic == nil {
		diagnostic = new(Diagnostic)
	}

	diagnostic.Errors = append(diagnostic.Errors, err)

	return diagnostic
}

func (diagnostic *Diagnostic) IsZero() bool {
	if diagnostic == nil {
		return true
	}

	return (len(diagnostic.Warnings) == 0) && (len(diagnostic.Errors) == 0)
}

func (diagnostic *Diagnostic) HasErrors() bool {
	return !diagnostic.IsZero() && len(diagnostic.Errors) > 0
}
