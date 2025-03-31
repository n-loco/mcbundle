package cli

var HelpTask = TaskDefs{
	Name:         "help",
	Aliases:      []string{"--help", "-h", "h", "/?", "-?"},
	Doc:          "prints this message.",
	Dependencies: 0,
	Execute: func(deps *EnvironmentDependencies) {

	},
}
