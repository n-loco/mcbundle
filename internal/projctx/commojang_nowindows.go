//go:build !windows

package projctx

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/txtui"
)

func warnComMojangPath(should bool, diagnostic alert.Diagnostic) {
	if should {
		var _, exists = os.LookupEnv(comMojangPathVarKey)
		if !exists {
			diagnostic.AppendWarning(comMojangAlert())
		}
	}
}

func comMojangPath(diagnostic alert.Diagnostic) (path string) {
	value, exists := os.LookupEnv(comMojangPathVarKey)
	if !exists {
		diagnostic.AppendError(comMojangAlert())
	}
	path = value
	return
}

func comMojangAlert() alert.Alert {
	return alert.AlertTF(
		"deducing the path to com.mojang is not possible", nil,
		"consider adding %s to your environment variables", []any{
			txtui.EscapeItalic + comMojangPathVarKey + txtui.EscapeReset,
		},
	)
}
