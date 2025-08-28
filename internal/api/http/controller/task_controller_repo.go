package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
	"net/http"
)

type taskController struct {
	taskUsecase usecase.CreateTask
}

type TaskController interface {
	Create(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTask(c *gin.Context)
}

func NewTaskController(tu usecase.CreateTask) TaskController {
	return &taskController{taskUsecase: tu}
}

func (tc *taskController) Create(c *gin.Context) {

	var taskRequest usecase.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\nError: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	newTask, err := tc.taskUsecase.Create(taskRequest)
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\n Error: %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	c.JSON(http.StatusCreated, newTask)
}
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
