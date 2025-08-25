package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/api/http/handler"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/domain"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
	"net/http"
)

type AppController struct {
	Task interface{ TaskController }
}

type taskController struct {
	taskHandler handler.TaskHandler
}

type TaskController interface {
	Create(c *gin.Context)
}

func NewTaskController(th handler.TaskHandler) TaskController {
	return &taskController{taskHandler: th}
}

// CreateUser create a user in the postgres database
func (tc *taskController) Create(c *gin.Context) {

	var taskRequest usecase.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\nError = %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	var service domain.TaskRepository
	handler := handler.NewTaskHandler(service)
	newTask, err := handler.Create(taskRequest)
	if err != nil {
		c.JSON(err.Code, err.Message)
	}
	c.JSON(http.StatusCreated, newTask)
}
