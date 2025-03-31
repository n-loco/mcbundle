package cli

type TaskDefs struct {
	Dependencies DependenciesFlags
	Name         string
	Aliases      []string
	Doc          string
	Execute      func(*EnvironmentDependencies)
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
	registerTask(&HelpTask)
}

func GetTask() TaskDefs {
	return TaskDefs{}
}
