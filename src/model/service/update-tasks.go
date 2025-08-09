package service

import (
	"errors"
	"fmt"

	"github.com/rodrigodip/toDo-API/src/config/database/mysql"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model"
	"gorm.io/gorm"
)

func UpdateTaskByID(id string, req request.TaskRequest) *rest_err.RestErr {

	taskMux.Lock()
	defer taskMux.Unlock()

	db, err := database.GetDB()
	if err != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return restError
	}

	var task model.TaskData
	if err := db.First(&task, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			restError := rest_err.NewNotFoundError(
				fmt.Sprintf("no task found with [id:%s]", id),
			)
			return restError
		}
		//if it isn't a NotFound it's a InternalError
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return restError
	}

	db.Model(&task).Select("Title", "Description").Updates(response.TaskResponse{Title: req.Title, Description: req.Description})

	return nil
}
