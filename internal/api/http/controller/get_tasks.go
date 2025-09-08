package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

// GetTasks retrieves all Tasks.
// @Summary Find all Tasks
// @Description Retrieves all Tasks details.
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {object} usecase.TaskDtoOutput "Tasks information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: No tasks found."
// @Failure 404 {object} rest_err.RestErr "Error: ID not found"
// @Router /allTasks [get]
func (tc *taskController) GetTasks(c *gin.Context) {
	allTasks, err := tc.taskUsecase.GetTasks()
	if err != nil {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintf("No tasks found.\n Error: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, allTasks)
}
