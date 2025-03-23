//go:build windows

package envdeps

import (
	"os"
	"strings"
)

func warnComMojangPath(b bool) {

}

func getComMojangPath() string {
	value, exists := os.LookupEnv("AUTOCRAFTER_COM_MOJANG_PATH")

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
