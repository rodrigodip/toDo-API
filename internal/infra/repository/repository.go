package repository

import "github.com/rodrigodip/toDo-API/internal/domain"

type task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}

type TaskRepository interface {
	Create(id, title, description string, completed bool) error
	GetTasks() ([]domain.Task, error)
	GetTask(id string) (domain.Task, error)
	UpdateTask(id string) (domain.Task, error)
	DeleteTask(id string) (domain.Task, error)
	SetTaskDone(id string) error
}
