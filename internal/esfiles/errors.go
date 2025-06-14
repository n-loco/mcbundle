package esfiles

import (
	"fmt"
)

type resolveNativeModErr struct {
	NativeModule string
}

func (nativeModErr resolveNativeModErr) Error() string {
	return fmt.Sprintf("unnable to resolve native module %s", nativeModErr.NativeModule)
}
