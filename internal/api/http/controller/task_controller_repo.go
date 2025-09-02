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

// DeleteTask deletes a task with the specified ID.
// @Summary Delete task
// @Description Deletes a task based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param taskId path string true "ID of the task to be deleted"
// @Success 200 "Task {id: %s} DELETED"
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /deleteTask/{taskId} [delete]
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

// UpdateTask updates Tasks information with the specified ID.
// @Summary Update Tasks
// @Description Updates Tasks details based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param TaskId path string true "ID of the Tasks to be updated"
// @Param TasksRequest body usecase.TaskDtoOutput true "Tasks information for update"
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

// SetTaskDone Defines a task as Completes using the specified ID.
// @Summary Mark a task as completed
// @Description Defines a task as completed based on the ID provided as a parameter.
// @Tags Tasks
// @Accept json
// @Produce json
// @Param taskId path string true "ID of the task to be completed"
// @Success 200 "Task {taskId} setted as Done"
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /setTaskDone/{taskId} [put]
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
