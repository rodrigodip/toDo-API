package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/model/service"
)

func GetTasksAll(c *gin.Context) {
	tasks := service.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}
