//go:build linux

package cli

import "os/exec"

func open(uri string) (err error) {
	var cmd = exec.Command("xdg-open", uri)

	err = cmd.Start()

	if err == nil {
		err = cmd.Wait()
	}

	return
}
