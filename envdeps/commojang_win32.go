//go:build windows

package envdeps

import (
	"os"
	"strings"
)

func warnComMojangPath(_ bool) {

}

func getComMojangPath() string {
	value, exists := os.LookupEnv(comMojangVarKey)

	if exists {
		return value
	}

	appDataDir := os.Getenv("LocalAppData")
	comMojangDir := strings.Join([]string{
		appDataDir,
		"Packages\\Microsoft.MinecraftUWP_8wekyb3d8bbwe\\LocalState\\games\\com.mojang",
	}, "\\")

	return strings.ReplaceAll(comMojangDir, "\\\\", "\\")
}
