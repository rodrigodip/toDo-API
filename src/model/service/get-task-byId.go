package service

import (
	"errors"
	"fmt"

	"github.com/rodrigodip/toDo-API/src/config/database/mysql"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model"
	"gorm.io/gorm"
)

func GetTaskByID(id string) (response.TaskResponse, *rest_err.RestErr) {

	db, err := database.GetDB()
	if err != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return response.TaskResponse{}, restError
	}

	var retrieved model.TaskData

	if err := db.First(&retrieved, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			restError := rest_err.NewNotFoundError(
				fmt.Sprintf("no task found with [id:%s]", id),
			)
			return response.TaskResponse{}, restError
		}
		//if it isn't a NotFound it's a InternalError
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return response.TaskResponse{}, restError
	}

	return response.TaskResponse{
		ID:          retrieved.ID,
		Title:       retrieved.Title,
		Description: retrieved.Description,
		Completed:   retrieved.Completed,
	}, nil
}
