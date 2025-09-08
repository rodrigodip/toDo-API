package usecase

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
)

func (ct *TaskRepository) GetTasks() ([]TaskDtoOutput, error) {
	var allTasks []domain.Task
	allTasks, err := ct.Repository.GetTasks()
	if err != nil {
		return []TaskDtoOutput{}, err
	}
	var output []TaskDtoOutput
	for _, tasks := range allTasks {
		output = append(output, TaskDtoOutput{
			ID:          tasks.ID,
			Title:       tasks.Title,
			Description: tasks.Description,
			Completed:   tasks.Completed,
		})
	}
	return output, nil
}
