package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

// Create Creates a new task
// @Summary Create a new task
// @Description Create a new task with the provided title and description
// @Tags Tasks
// @Accept json
// @Produce json
// @Param TaskRequest body usecase.CreateTaskRequest true "Task information for registration"
// @Success 200 {object} usecase.TaskDtoOutput
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /createTask [post]
func (tc *taskController) CreateTask(c *gin.Context) {

	var taskRequest usecase.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\nError: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	newTask, err := tc.taskUsecase.CreateTask(taskRequest)
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\n Error: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusCreated, newTask)
}
