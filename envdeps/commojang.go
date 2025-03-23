package envdeps

import "os"

func WarnComMojangPath(b bool) {
	// TODO
}

func GetComMojangPath() (string, bool) {
	return os.LookupEnv("AUTOCRAFTER_COM_MOJANG_PATH")
}
