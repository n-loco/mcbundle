//go:build darwin

package cli

import "os/exec"

func open(uri string) (err error) {
	var cmd = exec.Command("open", uri)

	err = cmd.Start()

	if err == nil {
		err = cmd.Wait()
	}

	return
}
