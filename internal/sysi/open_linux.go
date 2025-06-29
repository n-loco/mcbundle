//go:build linux

package sysi

import (
	"os/exec"
)

var isHasOpenSupportCached = false
var hasOpenSupportCache = false
var hasXdgCache = false

func HasOpenSupport() bool {
	if isHasOpenSupportCached {
		return hasOpenSupportCache
	}

	hasXdgCache = hasExec("xdg-open")

	var result = hasXdgCache

	if !result {
		result = result || hasExec("gio")
	}

	isHasOpenSupportCached = true
	hasOpenSupportCache = result
	return result
}

func Open(uri string) (err error) {
	if !HasOpenSupport() {
		panic(unsupportedOpenMessage)
	}

	var cmd *exec.Cmd

	if hasXdgCache {
		cmd = exec.Command("xdg-open", uri)
	} else {
		cmd = exec.Command("gio", "open", uri)
	}

	err = cmd.Start()

	if err == nil {
		err = cmd.Wait()
	}

	return
}

func hasExec(exeName string) bool {
	var _, err = exec.LookPath(exeName)
	return err == nil
}
