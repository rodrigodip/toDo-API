package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rodrigodip/toDo-API/src/controller/routes"
)

func main() {
	router := gin.Default()
	routes.InitGroup(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
