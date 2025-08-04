package service

import (
	"fmt"
	"strconv"

	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func UpdateTaskByID(id string, req request.TaskRequest) (response.TaskRespose, *rest_err.RestErr) {
	taskMux.Lock()
	defer taskMux.Unlock()

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return response.TaskRespose{}, restErr
	}

	for _, t := range repository.TaskRepository {

		if t.GetId() == intId {
			t.SetTitle(req.Title)
			t.SetDescription(req.Description)

			return response.TaskRespose{
				ID:          t.ID,
				Title:       t.Title,
				Description: t.Description,
				Completed:   t.Completed,
			}, nil
		}
	}
	restErr := rest_err.NewNotFoundError(
		fmt.Sprintln("Error: ID not Found"),
	)

	return response.TaskRespose{}, restErr
}
