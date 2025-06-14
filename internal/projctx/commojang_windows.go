//go:build windows

package projctx

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/internal/alert"
)

func warnComMojangPath(bool) (_ *alert.Diagnostic) {
	return
}

func comMojangPath() (path string, _ *alert.Diagnostic) {
	value, exists := os.LookupEnv(comMojangPathVarKey)

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
