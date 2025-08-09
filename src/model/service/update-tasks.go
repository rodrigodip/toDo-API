package service

import (
	"fmt"
	"strconv"

	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func UpdateTaskByID(id string, req request.TaskRequest) *rest_err.RestErr {
	taskMux.Lock()
	defer taskMux.Unlock()

	if len(repository.TaskRepository) < 1 {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintln("Error: No Tasks Found"),
		)
		return restErr
	}

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return restErr
	}

	for idx, t := range repository.TaskRepository {

		if t.GetId() == uint(intId) {
			repository.TaskRepository[idx].SetTitle(req.Title)
			repository.TaskRepository[idx].SetDescription(req.Description)

			return nil
		}
	}
	restErr := rest_err.NewNotFoundError(
		fmt.Sprintln("Error: ID not Found"),
	)

	return restErr
}
