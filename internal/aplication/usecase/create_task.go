package usecase

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
	IDgenerator "github.com/rodrigodip/toDo-API/pkg/id_generator"
)

type CreateTask struct {
	Repository domain.TaskRepository
}

func Newtask(repository domain.TaskRepository) *CreateTask {
	return &CreateTask{Repository: repository}
}

func (ct *CreateTask) Create(input CreateTaskRequest) (TaskDtoOutput, error) {
	newTask := domain.NewTask()
	newTask.Title = input.Title
	newTask.Description = input.Description
	err := newTask.TaskValidation()
	if err != nil {
		return TaskDtoOutput{}, err
	}
	output := TaskDtoOutput{
		ID:          IDgenerator.NewID(),
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   false,
	}
	err = ct.Repository.Create(output.ID, output.Title, output.Description, output.Completed)
	return output, nil
}
