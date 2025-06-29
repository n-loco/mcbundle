//go:build darwin

package sysi

import "os/exec"

func HasOpenSupport() bool {
	return true
}

func Open(uri string) (err error) {
	var cmd = exec.Command("open", uri)

	err = cmd.Start()

	if err == nil {
		err = cmd.Wait()
	}

	return
}
