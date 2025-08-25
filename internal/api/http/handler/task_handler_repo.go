package handler

import (
	"fmt"

	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/domain"
	"github.com/rodrigodip/toDo-API/internal/infra/db/mysql"
	"github.com/rodrigodip/toDo-API/internal/infra/repository"
)

type taskHandler struct {
	TaskRepository domain.TaskRepository
}

type TaskHandler interface {
	Create(input usecase.CreateTaskRequest) (usecase.TaskDtoOutput, error)
}

func NewTaskHandler(t domain.TaskRepository) TaskHandler {
	return &taskHandler{TaskRepository: t}
}

func (th *taskHandler) Create(input usecase.CreateTaskRequest) (usecase.TaskDtoOutput, error) {
	db, err := mysql.NewDataBaseConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	repo := repository.NewTaskRepositoryDB(db)
	service := usecase.Newtask(repo)
	req := usecase.CreateTaskRequest{

		Title:       input.Title,
		Description: input.Description,
	}
	var output usecase.TaskDtoOutput
	output, err = service.Create(req)
	if err != nil {
		return usecase.TaskDtoOutput{}, err
	}

	return output, nil

	// db, err := mysql.NewDataBaseConnection()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// repo := repository.NewTaskRepositoryDB(db)
	// service := usecase.Newtask(repo)
	// imput := usecase.CreateTaskRequest{
	// 	Title:       title,
	// 	Description: description,
	// }
	// _, err = service.Create(imput)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// return nil
}
