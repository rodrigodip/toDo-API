package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/model/service"
)

// SetTaskDone Defines a task as Completes using the specified ID.
// @Summary Mark a task as completed
// @Description Defines a task as completed based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param taskId path string true "ID of the task to be completed"
// @Success 200 "Task {taskId} setted as Done"
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /setTaskDone/{taskId} [put]
func SetTaskDone(c *gin.Context) {

	taskId := c.Param("id")

	err := service.SetTaskDone(taskId)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	sucessMessage := fmt.Sprintf("Task [id:%s] setted as Done", taskId)
	c.JSON(http.StatusOK, sucessMessage)
}
