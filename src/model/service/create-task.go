package service

import (
	"fmt"
	"sync"

	"github.com/rodrigodip/toDo-API/src/config/database/mysql"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model"
)

var taskMux sync.Mutex

func CreateTask(req request.TaskRequest) (response.TaskResponse, *rest_err.RestErr) {
	taskMux.Lock()
	defer taskMux.Unlock()

	db, err := database.GetDB()
	if err != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return response.TaskResponse{}, restError
	}

	newTask := model.TaskData{
		Title:       req.Title,
		Description: req.Description,
	}

	created := db.Create(&newTask)
	if err = created.Error; err != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return response.TaskResponse{}, restError
	}

	return response.TaskResponse{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   newTask.Completed,
	}, nil
}
