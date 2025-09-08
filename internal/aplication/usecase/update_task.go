package usecase

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
)

func (ct *TaskRepository) UpdateTask(id, title, description string) (TaskDtoOutput, error) {
	newTask := domain.NewTaskDomain()
	newTask.Title = title
	newTask.Description = description
	err := newTask.TaskValidation()
	if err != nil {
		return TaskDtoOutput{}, err
	}
	updated, err := ct.Repository.UpdateTask(id, newTask.Title, newTask.Description)
	if err != nil {
		return TaskDtoOutput{}, err
	}
	output := TaskDtoOutput{
		ID:          id,
		Title:       updated.Title,
		Description: updated.Description,
		Completed:   false,
	}

	return output, nil
}
