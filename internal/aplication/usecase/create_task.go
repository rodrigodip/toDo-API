package usecase

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
	"github.com/rodrigodip/toDo-API/pkg/id_generator"
)

type CreateTask struct {
	Repository domain.TaskRepository
}

func NewTaskRepository(repository domain.TaskRepository) *CreateTask {
	return &CreateTask{Repository: repository}
}

func (ct *CreateTask) Create(input CreateTaskRequest) (TaskDtoOutput, error) {
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
func (ct *CreateTask) GetTask(id string) (TaskDtoOutput, error) {
	var task domain.Task
	task, err := ct.Repository.GetTask(id)
	if err != nil {
		return TaskDtoOutput{}, nil
	}

	output := TaskDtoOutput{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
	}

	return output, nil
}
func (ct *CreateTask) UpdateTask(id, title, description string) (TaskDtoOutput, error) {
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
func (ct *CreateTask) DeleteTask(id string) error {
	err := ct.Repository.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}
func (ct *CreateTask) SetTaskDone(id string) error {
	err := ct.Repository.SetTaskDone(id)
	if err != nil {
		return err
	}
	return nil
}
