package usecase

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
)

func (ct *TaskRepository) GetTask(id string) (TaskDtoOutput, error) {
	var task domain.Task
	task, err := ct.Repository.GetTask(id)
	if err != nil {
		return TaskDtoOutput{}, err
	}

	output := TaskDtoOutput{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}

	return output, nil
}
