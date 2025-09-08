package usecase

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
	"github.com/rodrigodip/toDo-API/pkg/id_generator"
)

func (ct *TaskRepository) CreateTask(input CreateTaskRequest) (TaskDtoOutput, error) {
	newTask := domain.NewTaskDomain()
	newTask.Title = input.Title
	newTask.Description = input.Description
	err := newTask.TaskValidation()
	if err != nil {
		return TaskDtoOutput{}, err
	}
	generator := IDgenerator.NewTimestampIDGenerator()
	output := TaskDtoOutput{
		ID:          generator.NewID(),
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   false,
	}
	err = ct.Repository.CreateTask(output.ID, output.Title, output.Description, output.Completed)
	if err != nil {
		return TaskDtoOutput{}, err
	}
	return output, nil
}
