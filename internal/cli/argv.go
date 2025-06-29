package cli

import "os"

type argvIterator struct {
	argv []string
	pos  int
}

func newArgvIterator() *argvIterator {
	return &argvIterator{
		argv: os.Args[1:],
	}
}

func (argvI *argvIterator) consume() string {
	if argvI.pos > len(argvI.argv)-1 {
		panic("argv: out of bounds")
	}

	var v = argvI.argv[argvI.pos]
	argvI.pos++

	return v
}

func (argvI *argvIterator) remaining() int {
	return len(argvI.argv) - argvI.pos
}

func (argvI *argvIterator) hasNext() bool {
	return argvI.remaining() > 0
}
