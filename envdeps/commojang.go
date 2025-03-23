//go:build !windows

package envdeps

import (
	"os"

	"github.com/redrock/autocrafter/cli"
)

func warnComMojangPath(b bool) {
	_, exists := os.LookupEnv("AUTOCRAFTER_COM_MOJANG_PATH")
	if !exists {
		cli.Wprint("AUTOCRAFTER_COM_MOJANG_PATH variable not defined")
	}
}

func getComMojangPath() string {
	value, exists := os.LookupEnv("AUTOCRAFTER_COM_MOJANG_PATH")
	if !exists {
		os.Exit(1)
	}
	return value
}
