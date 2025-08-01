package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mcbundle/mcbundle/tools/glob"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var allFiles []string

	for _, arg := range os.Args[1:] {
		var files, err = glob.Glob(arg)

		if err != nil {
			return
		}

		allFiles = append(allFiles, files...)
	}

	fmt.Print(strings.Join(allFiles, " "))
}
