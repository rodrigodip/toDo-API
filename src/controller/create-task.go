package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/model"
	"github.com/rodrigodip/toDo-API/src/model/service"
	"net/http"
)

func CreateTask(c *gin.Context) {

	var taskRequest model.Task

	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintf("There are some incorrect fields.\nError = %s\n", err.Error()),
		)
		c.JSON(restErr.Code, restErr)
		return
	}
	created := service.CreateTask(taskRequest)
	c.JSON(http.StatusCreated, created)
}
