package domain

type TaskRepository interface {
	Create(id, title, description string, completed bool) error
	//GetTask
	//GetTasks
	//UpdateTask
	//DeleteTask
	//SetTaskDone
}
