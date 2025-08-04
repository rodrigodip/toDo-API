package service

import (
	"fmt"
	"strconv"

	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func DeleteTaskByID(id string) *rest_err.RestErr {
	taskMux.Lock()
	defer taskMux.Unlock()

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return restErr
	}

	for i, t := range repository.TaskRepository {
		if t.GetId() == intId {
			repository.DeleteTask(i)
			return nil
		}
	}
	restErr := rest_err.NewNotFoundError(
		fmt.Sprintln("Error: Task not Found"),
	)
	return restErr
}
