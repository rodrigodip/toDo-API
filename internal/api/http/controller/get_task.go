package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

// GetTaskById retrieves Tasks information based on the provided Tasks ID.
// @Summary Find Tasks by ID
// @Description Retrieves Tasks details based on the Tasks ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param tasksId path string true "ID of the Tasks to be retrieved"
// @Success 200 {object} usecase.TaskDtoOutput "Tasks information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: ID must be a number"
// @Failure 404 {object} rest_err.RestErr "Error: ID not found"
// @Router /taskById/{tasksId} [get]
func (tc *taskController) GetTask(c *gin.Context) {
	taskId := c.Param("id")
	task, err := tc.taskUsecase.GetTask(taskId)
	if err != nil {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintf("No tasks found.\n Error: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, task)
}
