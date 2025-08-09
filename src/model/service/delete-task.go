package service

import (
	"fmt"

	"github.com/rodrigodip/toDo-API/src/config/database/mysql"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/model"
)

func DeleteTaskByID(id string) *rest_err.RestErr {
	taskMux.Lock()
	defer taskMux.Unlock()

	db, err := database.GetDB()
	if err != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return restError
	}

	deleted := db.Unscoped().Delete(&model.TaskData{}, id)
	if deleted.Error != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", deleted.Error),
		)
		return restError
	}
	if deleted.RowsAffected == 0 {
		restError := rest_err.NewNotFoundError(
			fmt.Sprintf("no task found with [id:%s]", id),
		)
		return restError
	}

	return nil
}
