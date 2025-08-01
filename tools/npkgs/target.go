package main

import (
	"strings"
)

func splitTargetDouble(targetDouble string) (os string, cpu string) {
	var t = strings.Split(targetDouble, "-")
	os = t[0]
	cpu = t[1]
	return
}

func displayCPU(nodecpu string) (dcpu string) {
	switch nodecpu {
	case "x64":
		dcpu = "x86_64"
	case "ia32":
		dcpu = "x86"
	case "arm":
		dcpu = "ARM"
	case "arm64":
		dcpu = "ARM 64-bit"
	case "mips":
		dcpu = "MIPS (BE)"
	case "mipsel":
		dcpu = "MIPS (LE)"
	case "mips64":
		dcpu = "MIPS 64-bit (BE)"
	case "mips64el":
		dcpu = "MIPS 64-bit (LE)"
	case "ppc64":
		dcpu = "PowerPC 64-bit (BE)"
	case "ppc64le":
		dcpu = "PowerPC 64-bit (LE)"
	case "s390":
		dcpu = "IBM System/390"
	case "s390x":
		dcpu = "IBM System/390x"
	case "riscv64":
		dcpu = "RISC-V 64-bit"
	default: // generic fallback
		dcpu = nodecpu
	}
	return
}

func displayOS(nodeos string) (dos string) {
	switch nodeos {
	case "win32":
		dos = "Windows"
	case "darwin":
		dos = "macOS"
	case "freebsd":
		dos = "FreeBSD"
	case "openbsd":
		dos = "OpenBSD"
	case "sunos":
		dos = "Solaris"
	case "linux":
		fallthrough
	case "android":
		fallthrough
	default: // generic fallback
		dos = strings.ToTitle(nodeos[:1]) + nodeos[1:]
	}
	return
}
