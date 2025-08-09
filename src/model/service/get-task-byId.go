package service

import (
	"fmt"
	"strconv"

	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func GetTaskByID(id string) (response.TaskResponse, *rest_err.RestErr) {

	if len(repository.TaskRepository) < 1 {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintln("Error: No Tasks Found"),
		)
		return response.TaskResponse{}, restErr
	}

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return response.TaskResponse{}, restErr
	}

	for _, t := range repository.TaskRepository {
		if t.GetId() == uint(intId) {
			return response.TaskResponse{
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
	return response.TaskResponse{}, restErr
}
