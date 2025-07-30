package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/model"
	"github.com/rodrigodip/toDo-API/src/model/service"
)

func UpdateTask(c *gin.Context) {

	var taskToUpdate model.Task
	taskId := c.Param("id")

	if err := c.ShouldBindJSON(&taskToUpdate); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("Bad request body.\nError = %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}

	task, err := service.UpdateTaskByID(taskId, taskToUpdate)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(http.StatusOK, task)
}
