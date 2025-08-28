package domain

type TaskRepository interface {
	Create(id, title, description string, completed bool) error
	GetTasks() ([]Task, error)
	//GetTask
	//UpdateTask
	//DeleteTask
	//SetTaskDone
}
