package cli

import (
	"os"

	"github.com/n-loco/bpbuild/internal/projctx"
	"github.com/n-loco/bpbuild/internal/terminal"
)

type TaskDefs struct {
	Requires projctx.EnvRequireFlags
	Name     string
	Aliases  []string
	Doc      string
	Execute  func(*projctx.ProjectContext)
}

var taskMap = map[string]*TaskDefs{}
var taskList = []*TaskDefs{}

func registerTask(task *TaskDefs) {
	taskList = append(taskList, task)
	taskMap[task.Name] = task
	for _, alias := range task.Aliases {
		taskMap[alias] = task
	}
}

func SetupTasks() {
	registerTask(&devTask)
	registerTask(&packTask)
	registerTask(&buildTask)

	// special case
	registerTask(&helpTask)
	taskMap["-?"] = &helpTask
	taskMap["/?"] = &helpTask
	taskMap["h"] = &helpTask
}

func GetTask() *TaskDefs {
	if len(os.Args) < 2 {
		return &helpTask
	}

	taskName := os.Args[1]

	taskDefs, exists := taskMap[taskName]

	if !exists {
		terminal.Eprintf("unknown task: %s;\n", taskName)
		terminal.Eprint("use " + terminal.UnderlineWhite + "autocrafter help" + terminal.Reset + " to see a list of tasks.\n")
		os.Exit(1)
	}

	return taskDefs
}
