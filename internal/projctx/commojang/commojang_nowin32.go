//go:build !windows

package commojang

import (
	"os"

	"github.com/n-loco/bpbuild/internal/terminal"
)

func WarnComMojangPath(should bool) {
	if should {
		_, exists := os.LookupEnv(comMojangPathVarKey)
		if !exists {
			terminal.Wprint("deducing the path to com.mojang is not possible\n")
			terminal.Wprintf("consider adding %s to your environment variables\n", comMojangPathVarKey)
		}
	}
}

func ComMojangPath() string {
	value, exists := os.LookupEnv(comMojangPathVarKey)
	if !exists {
		terminal.Eprint("path to com.mojang directory not set\n")
		terminal.Eprintf("please add %s to your environment variables\n", comMojangPathVarKey)
		os.Exit(1)
	}
	return value
}
