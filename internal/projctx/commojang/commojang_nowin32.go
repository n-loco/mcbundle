//go:build !windows

package commojang

import (
	"os"

	"github.com/n-loco/bpbuild/internal/txtui"
)

func WarnComMojangPath(should bool) {
	if should {
		_, exists := os.LookupEnv(comMojangPathVarKey)
		if !exists {
			txtui.PrePrint(txtui.UIPartErr, txtui.WarnPrefix, "deducing the path to com.mojang is not possible\n")
			txtui.Printf(txtui.UIPartErr, "consider adding %s to your environment variables\n", comMojangPathVarKey)
		}
	}
}

func ComMojangPath() string {
	value, exists := os.LookupEnv(comMojangPathVarKey)
	if !exists {
		txtui.PrePrint(txtui.UIPartErr, txtui.ErrPrefix, "path to com.mojang directory not found\n")
		txtui.Printf(txtui.UIPartErr, "please add %s to your environment variables\n", comMojangPathVarKey)
		os.Exit(1)
	}
	return value
}
