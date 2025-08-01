package node2go

import (
	"fmt"
	"strings"
)

func GoTargetPairFromNodeTarget(nodeTarget string) (goos string, goarch string, err error) {
	var nodeTargetDouble = strings.Split(nodeTarget, "-")
	var nodeos = nodeTargetDouble[0]
	var nodecpu = nodeTargetDouble[1]

	goos, err = GOOSFromNodeOS(nodeos)
	if err != nil {
		return
	}

	goarch, err = GOARCHFromNodeCPU(nodecpu)
	if err != nil {
		goos = ""
	}

	return
}

func GOARCHFromNodeCPU(nodecpu string) (goarch string, err error) {
	switch nodecpu {
	case "x64":
		goarch = "amd64"
	case "ia32":
		goarch = "386"
	case "mipsel":
		goarch = "mipsle"
	case "mips64el":
		goarch = "mips64le"
	// ============================= \\
	case "arm":
		fallthrough
	case "arm64":
		fallthrough
	case "mips":
		fallthrough
	case "mips64":
		fallthrough
	case "ppc64":
		fallthrough
	case "ppc64le":
		fallthrough
	case "s390":
		fallthrough
	case "s390x":
		fallthrough
	case "riscv64":
		goarch = nodecpu
	// ============================= \\
	default:
		err = fmt.Errorf("unknown, unsupported or invalid node cpu: %s", nodecpu)
	}

	return
}

func GOOSFromNodeOS(nodeos string) (goos string, err error) {
	switch nodeos {
	case "sunos":
		goos = "solaris"
	case "win32":
		goos = "windows"
	// ============================= \\
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "openbsd":
		fallthrough
	case "linux":
		fallthrough
	case "android":
		goos = nodeos
	// ============================= \\
	default:
		err = fmt.Errorf("unknown, unsupported or invalid node os: %s", nodeos)
	}

	return
}
