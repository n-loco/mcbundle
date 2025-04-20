package cli

import "github.com/n-loco/bpbuild/internal/projctx"

type commandInfo struct {
	name    string
	aliases []string
	doc     string
}

type command interface {
	info() *commandInfo
	execute([]string)
}

type simpleCommand struct {
	commandInfo
	requirements projctx.EnvRequireFlags
	execCommand  func(*projctx.ProjectContext)
}

func (cmd *simpleCommand) info() *commandInfo {
	return &cmd.commandInfo
}

func (cmd *simpleCommand) execute([]string) {
	//optDefMap := make(map[string]*optionDefinitions[T])
	//
	//for _, optDef := range cmd.optionsDefinitions {
	//	optDefMap[optDef.name] = &optDef
	//
	//	for _, alias := range optDef.aliases {
	//		optDefMap[alias] = &optDef
	//	}
	//}

	projCtx := projctx.CreateProjectContext(cmd.requirements)
	cmd.execCommand(&projCtx)
}
