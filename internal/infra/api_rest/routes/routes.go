package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/internal/api/http/controller"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func InitGroup(r *gin.RouterGroup, app controller.TaskController) {

	r.POST("/createTask", app.Create)
	r.GET("/allTasks", app.GetTasks)
	// r.GET("/taskById/:id", controller.GetTaskById)
	// r.PUT("/updateTask/:id", controller.UpdateTask)
	// r.PUT("/setTaskDone/:id", controller.SetTaskDone)
	// r.DELETE("/deleteTask/:id", controller.DeleteTask)
	//
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
