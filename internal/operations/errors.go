package operations

import (
	"fmt"
)

type NativeModuleError struct {
	NativeModule string
}

func (nativeModErr NativeModuleError) Error() string {
	return fmt.Sprintf("unnable to find native module %s's version", nativeModErr.NativeModule)
}
