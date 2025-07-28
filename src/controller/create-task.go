package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/config"
	"github.com/rodrigodip/toDo-API/src/controller/request"
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
	fmt.Println(taskRequest)
}
