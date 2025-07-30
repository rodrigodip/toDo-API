package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/model/service"
	"net/http"
)

func GetTaskById(c *gin.Context) {

	taskId := c.Param("id")

	task, err := service.GetTaskByID(taskId)
	if err != nil {
		c.JSON(err.Code, err.Message)
		return
	}
	c.JSON(http.StatusOK, task)
}
