//go:build !windows

package envdeps

import (
	"os"

	"github.com/redrock/autocrafter/cli"
)

func warnComMojangPath(b bool) {
	_, exists := os.LookupEnv(comMojangVarKey)
	if !exists && b {
		cli.Wprint("deducing the path to com.mojang is not possible;\n")
		cli.Wprintf("consider adding %s to your environment variables.\n", comMojangVarKey)
	}
}

func getComMojangPath() string {
	value, exists := os.LookupEnv(comMojangVarKey)
	if !exists {
		cli.Eprint("getenv: path to com.mojang directory not set;\n")
		cli.Eprintf("getenv: please add %s to your environment variables.\n", comMojangVarKey)
		os.Exit(1)
	}
	return value
}
