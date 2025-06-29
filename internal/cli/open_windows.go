//go:build windows

package cli

import (
	"golang.org/x/sys/windows"
)

func open(uri string) (err error) {
	var wuri *uint16

	wuri, err = windows.UTF16PtrFromString(uri)
	if err != nil {
		return
	}

	err = windows.ShellExecute(0, nil, wuri, nil, nil, windows.SW_SHOWNORMAL)
	return
}
