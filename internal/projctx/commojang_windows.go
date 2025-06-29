//go:build windows

package projctx

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
)

func warnComMojangPath(bool, alert.Diagnostic) {}

func comMojangPath(alert.Diagnostic) (path string) {
	var value, exists = os.LookupEnv(comMojangPathVarKey)

	if exists {
		path = value
	} else {
		path = filepath.Join(
			os.Getenv("LocalAppData"),
			"Packages",
			"Microsoft.MinecraftUWP_8wekyb3d8bbwe",
			"LocalState",
			"games",
			"com.mojang",
		)
	}

	return
}
