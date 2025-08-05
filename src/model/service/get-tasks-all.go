package service

import (
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func GetAllTasks() []response.TaskResponse {
	var tasks []response.TaskResponse

	for _, t := range repository.TaskRepository {
		tasks = append(tasks, response.TaskResponse(t.GetAll()))
	}
	return tasks
}
