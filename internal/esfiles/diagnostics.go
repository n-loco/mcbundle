package esfiles

import (
	"fmt"
	"strings"

	esbuild "github.com/evanw/esbuild/pkg/api"
)

type esbuildWrapperAlert esbuild.Message

func (alertW *esbuildWrapperAlert) Display() string {
	file := alertW.Location.File
	line := alertW.Location.Line
	column := alertW.Location.Column
	msg := strings.ToLower(alertW.Text[:1]) + alertW.Text[1:]

	return fmt.Sprintf("%s:%d:%d: %s", file, line, column, msg)
}

func (alertW *esbuildWrapperAlert) Tip() string {
	return ""
}

type entryFileNotFoundErrAlert struct {
	ExpectedFiles   []string
	BundlingContext BundlingContext
}

func (errAlert entryFileNotFoundErrAlert) Display() string {
	var ctxName = errAlert.BundlingContext.String()

	return "no entry file found for " + ctxName
}

func (errAlert entryFileNotFoundErrAlert) Tip() string {
	var fileCount = len(errAlert.ExpectedFiles)
	var ctxName = errAlert.BundlingContext.String()
	switch fileCount {
	case 0:
		panic("there is no possible entries for " + ctxName + " anyway")
	case 1:
		return "the entry file for the " + ctxName + " must be " + errAlert.ExpectedFiles[0]
	default:
		return "the entry file for the " + ctxName + " must be one of these: " + strings.Join(errAlert.ExpectedFiles, ", ")
	}
}

type tooManyEntriesError struct {
	BundlingContext BundlingContext
	ExpectedFiles   []string
}

func (errA tooManyEntriesError) Display() string {
	var ctxName = errA.BundlingContext.String()

	return "too many entry files found for the " + ctxName
}

func (errA tooManyEntriesError) Tip() string {
	var ctxName = errA.BundlingContext.String()
	return "the entry file for the " + ctxName + " must be only one of these: " + strings.Join(errA.ExpectedFiles, ", ")
}
