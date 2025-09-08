package usecase

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
)

type TaskRepository struct {
	Repository domain.TaskRepository
}

func NewTaskRepository(repository domain.TaskRepository) *TaskRepository {
	return &TaskRepository{Repository: repository}
}
