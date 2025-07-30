package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/controller"
)

func InitGroup(r *gin.RouterGroup) {
	r.POST("/createTask", controller.CreateTask)
	r.GET("/allTasks", controller.GetTasksAll)
	r.GET("/taskById/:id", controller.GetTaskById)
	r.PUT("/editTask/:id", controller.UpdateTask)
	r.DELETE("/deleteTask/:taskId", controller.DeleteTask)
}
