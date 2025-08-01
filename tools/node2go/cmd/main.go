package main

import (
	"fmt"
	"os"

	"github.com/mcbundle/mcbundle/tools/node2go"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	var goos, goarch, err = node2go.GoTargetPairFromNodeTarget(os.Args[1])

	if err == nil {
		fmt.Print(goos, " ", goarch)
	}
}
