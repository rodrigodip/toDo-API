package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/controller"
)

func InitGroup(r *gin.RouterGroup) {
	r.POST("/createTask", controller.CreateTask)
	r.GET("/allTasks", controller.GetTasksAll)
	r.GET("/tasksById/:taskId", controller.GetTaskById)
	r.PUT("/editTask/:taskId", controller.UpdateTask)
	r.DELETE("/deleteTask/:taskId", controller.DeleteTask)
}
