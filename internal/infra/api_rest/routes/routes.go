package routes

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitGroup(r *gin.RouterGroup) {
	r.POST("/createTask", controller.CreateTask)
	r.GET("/allTasks", controller.GetTasksAll)
	r.GET("/taskById/:id", controller.GetTaskById)
	r.PUT("/updateTask/:id", controller.UpdateTask)
	r.PUT("/setTaskDone/:id", controller.SetTaskDone)
	r.DELETE("/deleteTask/:id", controller.DeleteTask)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
