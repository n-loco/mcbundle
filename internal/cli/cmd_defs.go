package cli

import "github.com/n-loco/bpbuild/internal/projctx"

type commandDefinitions struct {
	name         string
	aliases      []string
	doc          string
	requirements projctx.EnvRequireFlags
	execCommand  func(*projctx.ProjectContext)
}

func (cmd *commandDefinitions) execute([]string) {
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
