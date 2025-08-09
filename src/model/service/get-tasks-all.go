package service

import (
	"fmt"

	"github.com/rodrigodip/toDo-API/src/config/database/mysql"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model"
)

func GetAllTasks() ([]response.TaskResponse, *rest_err.RestErr) {

	db, err := database.GetDB()
	if err != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return []response.TaskResponse{}, restError
	}

	var tasks []model.TaskData

	retrieved := db.Select("id, title, description, completed").Find(&tasks)
	if err = retrieved.Error; err != nil {
		restError := rest_err.NewInternalServerError(
			fmt.Sprintf("DB error: %s", err),
		)
		return []response.TaskResponse{}, restError
	}
	var allTasks []response.TaskResponse
	if len(tasks) < 1 {
		restError := rest_err.NewNotFoundError(
			fmt.Sprintln("No Tasks Found"),
		)
		return []response.TaskResponse{}, restError

	}
	for _, t := range tasks {
		allTasks = append(allTasks, response.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
		})
	}
	return allTasks, nil
}
