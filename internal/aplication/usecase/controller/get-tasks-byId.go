package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/model/service"
	"net/http"
)

// GetTaskById retrieves Tasks information based on the provided Tasks ID.
// @Summary Find Tasks by ID
// @Description Retrieves Tasks details based on the Tasks ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param tasksId path string true "ID of the Tasks to be retrieved"
// @Success 200 {object} response.TaskResponse "Tasks information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: ID must be a number"
// @Failure 404 {object} rest_err.RestErr "Error: ID not found"
// @Router /taskById/{tasksId} [get]
func GetTaskById(c *gin.Context) {

	taskId := c.Param("id")

	task, err := service.GetTaskByID(taskId)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(http.StatusOK, task)
}
