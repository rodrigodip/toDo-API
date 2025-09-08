package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
)

type taskController struct {
	taskUsecase usecase.TaskRepository
}

type TaskController interface {
	CreateTask(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	SetTaskDone(c *gin.Context)
}

func NewTaskController(tu usecase.TaskRepository) TaskController {
	return &taskController{taskUsecase: tu}
}
