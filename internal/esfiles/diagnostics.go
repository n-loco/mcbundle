package esfiles

import (
	"fmt"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
	"github.com/mcbundle/mcbundle/internal/alert"
)

type esbuildWrapperAlert esbuild.Message

func (alertW *esbuildWrapperAlert) Display() alert.AlertDisplay {
	var file = alertW.Location.File
	var line = alertW.Location.Line
	var column = alertW.Location.Column
	var msg = strings.ToLower(alertW.Text[:1]) + alertW.Text[1:]

	var message = fmt.Sprintf("%s:%d:%d: %s", file, line, column, msg)

	return alert.AlertDisplay{
		Message: message,
	}
}

type entryFileNotFoundErrAlert struct {
	ExpectedFiles   []string
	BundlingContext BundlingContext
}

func (errAlert entryFileNotFoundErrAlert) Display() alert.AlertDisplay {
	var ctxName = errAlert.BundlingContext.String()
	var message = "no entry file found for " + ctxName
	var tip string

	var fileCount = len(errAlert.ExpectedFiles)

	switch fileCount {
	case 0:
		panic("there is no possible entries for " + ctxName + " anyway")
	case 1:
		tip = "the entry file for the " + ctxName + " must be " + errAlert.ExpectedFiles[0]
	default:
		tip = "the entry file for the " + ctxName + " must be one of these: " + strings.Join(errAlert.ExpectedFiles, ", ")
	}

	return alert.AlertDisplay{
		Message: message,
		Tip:     tip,
	}
}

type tooManyEntriesError struct {
	BundlingContext BundlingContext
	ExpectedFiles   []string
}

func (errA tooManyEntriesError) Display() alert.AlertDisplay {
	var ctxName = errA.BundlingContext.String()
	var message = "too many entry files found for the " + ctxName
	var tip = "the entry file for the " + ctxName + " must be only one of these: " + strings.Join(errA.ExpectedFiles, ", ")

	return alert.AlertDisplay{
		Message: message,
		Tip:     tip,
	}
}
