package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/model/service"
)

// UpdateTask updates Tasks information with the specified ID.
// @Summary Update Tasks
// @Description Updates Tasks details based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param TaskId path string true "ID of the Tasks to be updated"
// @Param TasksRequest body request.TaskRequest true "Tasks information for update"
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateTask/{TaskId} [put]
func UpdateTask(c *gin.Context) {

	var taskToUpdate request.TaskRequest
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
