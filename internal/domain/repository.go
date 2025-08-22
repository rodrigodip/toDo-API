package domain

type TaskRepository interface {
	insert(id, title, description string, completed bool) error
}
