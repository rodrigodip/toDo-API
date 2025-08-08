package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/rodrigodip/toDo-API/docs"
	"github.com/rodrigodip/toDo-API/src/config/database/mysql"
	"github.com/rodrigodip/toDo-API/src/controller/routes"
)

// @title toDo-API
// @version 1.0
// @description API for crud operations on tasks
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {

	database.NewDataBaseConnection()

	router := gin.Default()
	routes.InitGroup(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
