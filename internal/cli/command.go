package cli

type empty = struct{}

type commandInfo struct {
	name    string
	aliases []string
	doc     string
	options []option
}

type command interface {
	info() *commandInfo
	execute([]string)
}

type optionInfo struct {
	name    string
	aliases []string
}

type option interface {
	info() *optionInfo
}
