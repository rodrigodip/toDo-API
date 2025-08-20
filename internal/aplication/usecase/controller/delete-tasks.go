package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/model/service"
	"net/http"
)

// DeleteTask deletes a task with the specified ID.
// @Summary Delete task
// @Description Deletes a task based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param taskId path string true "ID of the task to be deleted"
// @Success 200 "Task {taskId} was Deleted"
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteTask/{taskId} [delete]
func DeleteTask(c *gin.Context) {

	taskId := c.Param("id")

	err := service.DeleteTaskByID(taskId)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	sucessMessage := fmt.Sprintf("Task [id:%s] was Deleted", taskId)
	c.JSON(http.StatusOK, sucessMessage)
}
