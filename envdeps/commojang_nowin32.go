//go:build !windows

package envdeps

import (
	"os"

	"github.com/redrock/autocrafter/cli"
)

func warnComMojangPath(b bool) {
	_, exists := os.LookupEnv(comMojangVarKey)
	if !exists && b {
		cli.Wprint("AUTOCRAFTER_COM_MOJANG_PATH variable not defined\n")
	}
}

func getComMojangPath() string {
	value, exists := os.LookupEnv(comMojangVarKey)
	if !exists {
		os.Exit(1)
	}
	return value
}
