//go:build windows

package envdeps

import (
	"os"
	"path/filepath"
)

func warnComMojangPath(bool) {

}

func ComMojangPath() string {
	value, exists := os.LookupEnv(ComMojangPathVarKey)

	if exists {
		return value
	}

	return filepath.Join(
		os.Getenv("LocalAppData"),
		"Packages",
		"Microsoft.MinecraftUWP_8wekyb3d8bbwe",
		"LocalState",
		"games",
		"com.mojang",
	)
}
