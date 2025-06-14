//go:build !windows

package projctx

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/txtui"
)

func warnComMojangPath(should bool) (diagnostic *alert.Diagnostic) {
	if should {
		_, exists := os.LookupEnv(comMojangPathVarKey)
		if !exists {
			diagnostic = diagnostic.AppendWarning(comMojangAlert())
		}
	}
	return
}

func comMojangPath() (path string, diagnostic *alert.Diagnostic) {
	value, exists := os.LookupEnv(comMojangPathVarKey)
	if !exists {
		diagnostic = diagnostic.AppendError(comMojangAlert())
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
