package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/rodrigodip/toDo-API/docs"
	"github.com/rodrigodip/toDo-API/src/config/database/mysql"
	"github.com/rodrigodip/toDo-API/src/controller/routes"
	"github.com/rodrigodip/toDo-API/src/model"
)

// @title toDo-API
// @version 1.0
// @description API for crud operations on tasks
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {

	if err := database.NewDataBaseConnection(); err != nil {
		panic("error opening connection")
	}
	db, err := database.GetDB()
	if err != nil {
		panic("error getting connection")
	}
	if err := db.AutoMigrate(&model.TaskData{}); err != nil {
		panic("erro migrating")
	}

	router := gin.Default()
	routes.InitGroup(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
