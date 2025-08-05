package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/model/service"
)

// CreateTask Creates a new task
// @Summary Create a new task
// @Description Create a new task with the provided title and description
// @Tags Tasks
// @Accept json
// @Produce json
// @Param TaskRequest body request.TaskRequest true "Task information for registration"
// @Success 200 {object} response.TaskResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /createTask [post]
func CreateTask(c *gin.Context) {

	var taskRequest request.TaskRequest

	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\nError = %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	newTask := service.CreateTask(taskRequest)
	c.JSON(http.StatusCreated, newTask)
}
