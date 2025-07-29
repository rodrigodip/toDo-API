package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/request"
	"github.com/rodrigodip/toDo-API/src/model/repository"
)

func CreateTask(c *gin.Context) {

	var taskRequest request.TaskRequest

	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\nError = %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	repo := repository.TaskRepo

	fmt.Println(repo)
}
