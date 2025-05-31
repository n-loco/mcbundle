//go:build !windows

package commojang

import (
	"os"

	"github.com/mcbundle/mcbundle/internal/alert"
	"github.com/mcbundle/mcbundle/internal/txtui"
)

type ComMojangWarnAndErrAlert struct{}

func (ComMojangWarnAndErrAlert) Display() string {
	return "deducing the path to com.mojang is not possible"
}
func (ComMojangWarnAndErrAlert) Tip() string {
	return "consider adding " + txtui.EscapeItalic + comMojangPathVarKey + txtui.EscapeReset + " to your environment variables"
}

func WarnComMojangPath(should bool) (diagnostic *alert.Diagnostic) {
	if should {
		_, exists := os.LookupEnv(comMojangPathVarKey)
		if !exists {
			diagnostic = diagnostic.AppendWarning(&ComMojangWarnAndErrAlert{})
		}
	}
	return
}

func ComMojangPath() (path string, diagnostic *alert.Diagnostic) {
	value, exists := os.LookupEnv(comMojangPathVarKey)
	if !exists {
		diagnostic = diagnostic.AppendError(&ComMojangWarnAndErrAlert{})
	}
	path = value
	return
}
