package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

// UpdateTask updates Tasks information with the specified ID.
// @Summary Update Tasks
// @Description Updates Tasks details based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param TaskId path string true "ID of the Tasks to be updated"
// @Param TasksRequest body usecase.CreateTaskRequest true "Tasks information for update"
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /updateTask/{TaskId} [put]
func (tc *taskController) UpdateTask(c *gin.Context) {
	taskId := c.Param("id")
	var dataToUpdate usecase.CreateTaskRequest
	if err := c.ShouldBindJSON(&dataToUpdate); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("Bad request body.\n Error: %s", err.Error()),
		)
		c.JSON(restErr.Code, restErr.Message)
	}
	task, err := tc.taskUsecase.UpdateTask(taskId, dataToUpdate.Title, dataToUpdate.Description)
	if err != nil {
		restErr := rest_err.NewNotFoundError(
			fmt.Sprintf("No tasks found.\n Error: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusOK, task)
}
