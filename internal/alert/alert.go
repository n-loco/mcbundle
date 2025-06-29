package alert

import "fmt"

type Alert interface {
	Display() AlertDisplay
}

func AlertF(format string, a ...any) Alert {
	return &genericAlert{
		AlertDisplay: AlertDisplay{
			Message: fmt.Sprintf(format, a...),
		},
	}
}

func AlertTF(msgFormat string, msgA []any, tipFormat string, tipA []any) Alert {
	return &genericAlert{
		AlertDisplay: AlertDisplay{
			Message: fmt.Sprintf(msgFormat, msgA...),
			Tip:     fmt.Sprintf(tipFormat, tipA...),
		},
	}
}

type AlertDisplay struct {
	Message string
	Tip     string
}

type genericAlert struct {
	AlertDisplay
}

func (gAlert genericAlert) Display() AlertDisplay {
	return AlertDisplay{
		Message: gAlert.Message,
		Tip:     gAlert.Tip,
	}
}
