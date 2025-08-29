package domain

type TaskRepository interface {
	Create(id, title, description string, completed bool) error
	GetTasks() ([]Task, error)
	GetTask(id string) (Task, error)
	UpdateTask(id, title, description string) (Task, error)
	DeleteTask(id string) error
	SetTaskDone(id string) error
}
