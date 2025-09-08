package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
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
func (tc *taskController) SetTaskDone(c *gin.Context) {
	taskId := c.Param("id")
	err := tc.taskUsecase.SetTaskDone(taskId)
	if err != nil {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintf("No tasks found.\n Error: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	response := fmt.Sprintf("Task {id:%s} setted as DONE", taskId)
	c.JSON(http.StatusOK, response)
}
