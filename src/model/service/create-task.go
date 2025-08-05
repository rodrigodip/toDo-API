package service

import (
	"sync"

	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

var (
	taskMux sync.Mutex
	nextID  = 1
)

func CreateTask(req request.TaskRequest) response.TaskResponse {
	taskMux.Lock()
	defer taskMux.Unlock()

	var newTask model.TaskData

	newTask.SetId(nextID)
	newTask.SetTitle(req.Title)
	newTask.SetDescription(req.Description)

	nextID++
	repository.TaskRepository = append(repository.TaskRepository, newTask)

	return response.TaskResponse{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   newTask.Completed,
	}
}
