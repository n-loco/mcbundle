package cli

type CLIAction interface {
	Execute()
}

func GetAction() CLIAction {
	return nil
}
