//go:build !windows

package envdeps

import (
	"os"

	"github.com/redrock/autocrafter/cli"
)

func warnComMojangPath(should bool) {
	if should {
		_, exists := os.LookupEnv(ComMojangPathVarKey)
		if !exists {
			cli.Wprint("deducing the path to com.mojang is not possible;\n")
			cli.Wprintf("consider adding %s to your environment variables.\n", ComMojangPathVarKey)
		}
	}
}

func ComMojangPath() string {
	value, exists := os.LookupEnv(ComMojangPathVarKey)
	if !exists {
		cli.Eprint("path to com.mojang directory not set;\n")
		cli.Eprintf("please add %s to your environment variables.\n", ComMojangPathVarKey)
		os.Exit(1)
	}
	return value
}
