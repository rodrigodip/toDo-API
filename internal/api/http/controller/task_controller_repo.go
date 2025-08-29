package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

type taskController struct {
	taskUsecase usecase.CreateTask
}

type TaskController interface {
	Create(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	SetTaskDone(c *gin.Context)
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
