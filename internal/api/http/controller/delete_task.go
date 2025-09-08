package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

// DeleteTask deletes a task with the specified ID.
// @Summary Delete task
// @Description Deletes a task based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param taskId path string true "ID of the task to be deleted"
// @Success 200 "Task {id: %s} DELETED"
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteTask/{taskId} [delete]
func (tc *taskController) DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	err := tc.taskUsecase.DeleteTask(taskId)
	if err != nil {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintf("No tasks found.\n Error: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	response := fmt.Sprintf("Task {id: %s} DELETED", taskId)
	c.JSON(http.StatusOK, response)

}
