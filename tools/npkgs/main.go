package main

import (
	"os"
	"path/filepath"

	"github.com/mcbundle/mcbundle/assets"
)

func main() {
	if len(os.Args) < 3 {
		return
	}

	var baseDir = filepath.FromSlash(os.Args[1])

	for _, targetDouble := range os.Args[2:] {
		var nodeos, nodecpu = splitTargetDouble(targetDouble)

		var dos = displayOS(nodeos)
		var dcpu = displayCPU(nodecpu)

		var exename = ""

		if nodeos == "win32" {
			exename = ".exe"
		}

		var vars = variables{intern: map[string]string{
			"program_version": assets.ProgramVersion,
			"target_double":   targetDouble,
			"os":              nodeos,
			"cpu":             nodecpu,
			"display_os":      dos,
			"display_cpu":     dcpu,
			"exename":         exename,
		}}

		var finalPackage = vars.apply(templatePackage)
		var finalPublishPackage = vars.apply(templateReleasePackage)

		var packageDir = filepath.Join(baseDir, targetDouble)

		os.MkdirAll(packageDir, os.ModePerm)

		var packagePath = filepath.Join(packageDir, "package.json")
		var packagePublishPath = filepath.Join(packageDir, "package.publish.json")
		var thirdPartyPath = filepath.Join(packageDir, "THIRD_PARTY.md")

		var packageFile, pkgFErr = os.Create(packagePath)
		if pkgFErr != nil {
			return
		}
		defer packageFile.Close()

		var packagePublishFile, pkgrFErr = os.Create(packagePublishPath)
		if pkgrFErr != nil {
			return
		}
		defer packagePublishFile.Close()

		var thirdPartyFile, tpfErr = os.Create(thirdPartyPath)
		if tpfErr != nil {
			return
		}
		defer thirdPartyFile.Close()

		packageFile.Write([]byte(finalPackage))
		packagePublishFile.Write([]byte(finalPublishPackage))
		thirdPartyFile.Write([]byte(thirdParty))
	}
}
