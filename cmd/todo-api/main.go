package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrigodip/toDo-API/internal/infra/api_rest/routes"
	"github.com/rodrigodip/toDo-API/internal/infra/db/mysql"
	"github.com/rodrigodip/toDo-API/pkg/dependencies"
)

// @title toDo-API
// @version 1.0
// @description API for crud operations on tasks
// @host localhost:8080
// @BasePath /
// @schemes http
// @
func main() {
	godotenv.Load()
	router := gin.Default()
	database, err := mysql.NewDataBaseConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	userController, _ := dependencies.Init(database)
	routes.InitGroup(&router.RouterGroup, userController)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
