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
	if err != nil {
		return TaskDtoOutput{}, err
	}
	return output, nil
}
func (ct *CreateTask) GetTasks() ([]TaskDtoOutput, error) {
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
