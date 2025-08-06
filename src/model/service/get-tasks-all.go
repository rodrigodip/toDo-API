package service

import (
	"fmt"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func GetAllTasks() ([]response.TaskResponse, *rest_err.RestErr) {

	if len(repository.TaskRepository) < 1 {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintln("Error: No Tasks Found"),
		)
		return []response.TaskResponse{}, restErr
	}

	var tasks []response.TaskResponse

	for _, t := range repository.TaskRepository {
		tasks = append(tasks, response.TaskResponse(t.GetAll()))
	}
	return tasks, nil
}
