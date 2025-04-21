package cli

import (
	"github.com/n-loco/bpbuild/internal/alert"
	"github.com/n-loco/bpbuild/internal/projctx"
)

type UnknownOptionWarnAlert struct {
	OptionName string
}

func (warnAlert UnknownOptionWarnAlert) Display() string {
	return "ignored unknown option: " + warnAlert.OptionName
}

func (warnAlert UnknownOptionWarnAlert) Tip() string {
	return ""
}

type operationCommand[T any] struct {
	commandInfo
	optionMap    map[string]*operationOption[T]
	requirements projctx.EnvRequireFlags
	execCommand  func(*T, *projctx.ProjectContext) *alert.Diagnostic
}

func createOperationCommand[T any](
	cmdInfo commandInfo,
	requirements projctx.EnvRequireFlags,
	execCommand func(*T, *projctx.ProjectContext) *alert.Diagnostic,
	options []*operationOption[T],
) operationCommand[T] {
	cmdInfo.options = make([]option, 0, len(options))
	optMap := make(map[string]*operationOption[T])

	for _, opt := range options {
		cmdInfo.options = append(cmdInfo.options, opt)
		optMap[opt.name] = opt
		for _, alias := range opt.aliases {
			optMap[alias] = opt
		}
	}

	return operationCommand[T]{
		commandInfo:  cmdInfo,
		requirements: requirements,
		optionMap:    optMap,
		execCommand:  execCommand,
	}
}

func (cmd *operationCommand[T]) info() *commandInfo {
	return &cmd.commandInfo
}

func (cmd *operationCommand[T]) execute(optList []string) (diagnostic *alert.Diagnostic) {
	var o T

	for i := 0; i < len(optList); i++ {
		optName := optList[i]
		opt, ok := cmd.optionMap[optName]

		if ok {
			optSlice := optList[i+1:]
			i += opt.process(&o, optSlice)
		} else {
			diagnostic = diagnostic.AppendWarning(&UnknownOptionWarnAlert{OptionName: optName})
		}
	}

	projCtx, createCtxDiag := projctx.CreateProjectContext(cmd.requirements)

	diagnostic = diagnostic.Append(createCtxDiag)

	if diagnostic.HasErrors() {
		return
	}

	diagnostic = diagnostic.Append(cmd.execCommand(&o, &projCtx))

	return
}

type operationOption[T any] struct {
	optionInfo
	process func(*T, []string) int
}

func (opt *operationOption[T]) info() *optionInfo {
	return &opt.optionInfo
}
