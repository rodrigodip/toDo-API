package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/api/http/handler"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/domain"
	rest_err "github.com/rodrigodip/toDo-API/pkg/errors/rest-err"
)

type AppController struct {
	User interface{ UserController }
}

type userController struct {
	userInteractor domain.TaskRepository
}

type UserController interface {
	CreateUser(c *gin.Context)
}

func NewUserController(us domain.TaskRepository) UserController {
	return &userController{us}
}

// CreateUser create a user in the postgres database
func (uc *userController) CreateUser(c *gin.Context) {

	var taskRequest usecase.CreateTaskRequest
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\nError = %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	db, err := mysql.NewDataBaseConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	handler := handler.NewTaskHandler(taskRequest)
	if err != nil {
		c.JSON(err.Code, err.Message)
	}
	c.JSON(http.StatusCreated, newTask)
}
