package main

import "os"

func main() {
	if len(os.Args) < 2 {
		return
	}

	for _, path := range os.Args[1:] {
		os.RemoveAll(path)
	}
}
