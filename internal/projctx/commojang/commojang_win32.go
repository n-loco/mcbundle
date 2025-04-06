//go:build windows

package commojang

import (
	"os"
	"path/filepath"
)

func WarnComMojangPath(bool) {

}

func ComMojangPath() string {
	value, exists := os.LookupEnv(comMojangPathVarKey)

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
