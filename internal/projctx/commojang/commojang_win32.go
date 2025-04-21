//go:build windows

package commojang

import (
	"os"
	"path/filepath"

	"github.com/n-loco/bpbuild/internal/alert"
)

type ComMojangWarnAndErrAlert struct{}

func (ComMojangWarnAndErrAlert) Display() string {
	return ""
}

func (ComMojangWarnAndErrAlert) Tip() string {
	return ""
}

func WarnComMojangPath(bool) (_ *alert.Diagnostic) {
	return
}

func ComMojangPath() (path string, _ *alert.Diagnostic) {
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
