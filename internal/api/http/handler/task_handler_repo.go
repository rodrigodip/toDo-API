package handler

import (
	"fmt"

	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/domain"
	"github.com/rodrigodip/toDo-API/internal/infra/db/mysql"
	"github.com/rodrigodip/toDo-API/internal/infra/repository"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

type taskHandler struct {
	TaskRepository domain.TaskRepository
}

type TaskHandler interface {
	Create(input usecase.CreateTaskRequest) (usecase.TaskDtoOutput, *rest_err.RestErr)
}

func NewTaskHandler(t domain.TaskRepository) TaskHandler {
	return &taskHandler{TaskRepository: t}
}

func (th *taskHandler) Create(input usecase.CreateTaskRequest) (usecase.TaskDtoOutput, *rest_err.RestErr) {
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
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)

		return usecase.TaskDtoOutput{}, restError
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
