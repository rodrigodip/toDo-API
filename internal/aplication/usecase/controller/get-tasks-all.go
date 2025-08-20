package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/model/service"
)

// GetTasksAll retrieves all Tasks.
// @Summary Find all Tasks
// @Description Retrieves all Tasks details.
// @Tags Tasks
// @Accept json
// @Produce json
// @Success 200 {object} response.TaskResponse "Tasks information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: ID must be a number"
// @Failure 404 {object} rest_err.RestErr "Error: ID not found"
// @Router /allTasks [get]
func GetTasksAll(c *gin.Context) {
	tasks, err := service.GetAllTasks()
	if err != nil {
		c.JSON(err.Code, err.Message)
	}

	c.JSON(http.StatusOK, tasks)
}
