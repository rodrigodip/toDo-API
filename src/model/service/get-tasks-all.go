package service

import (
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func GetAllTasks() []response.TaskRespose {
	var tasks []response.TaskRespose

	for _, t := range repository.TaskRepository {
		tasks = append(tasks, response.TaskRespose(t.GetAll()))
	}
	return tasks
}
