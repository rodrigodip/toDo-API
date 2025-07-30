package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/model/service"
)

func SetTaskDone(c *gin.Context) {

	taskId := c.Param("id")

	err := service.SetTaskDone(taskId)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	sucessMessage := fmt.Sprintf("Task [id:%s] setted to Done", taskId)
	c.JSON(http.StatusOK, sucessMessage)
}
