package cli

import (
	"github.com/n-loco/bpbuild/internal/assets"
	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/txtui"
)

var versionTask = TaskDefs{
	Requires: 0,
	Name:     "version",
	Aliases:  []string{"--version", "-v"},
	Doc:      "prints bpbuild's version.",
	Execute: func(*projctx.ProjectContext) {
		txtui.Printf(txtui.UIPartOut, "%s\n", assets.ProgramVersion)
	},
}
