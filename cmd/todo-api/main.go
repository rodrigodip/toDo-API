package main

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rodrigodip/toDo-API/internal/api/http/handler"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/domain"
)

// @title toDo-API
// @version 1.0
// @description API for crud operations on tasks
// @host localhost:8080
// @BasePath /
// @schemes http
// @license MIT
func main() {
	godotenv.Load()
	imput := usecase.CreateTaskRequest{
		Title:       "teste handler",
		Description: "handler funciona!",
	}
	var service domain.TaskRepository
	handler := handler.NewTaskHandler(service)
	newTesk, err := handler.Create(imput)
	if err != nil {
		fmt.Println(err.Error())
	}
	output, err := json.Marshal(newTesk)
	if err != nil {

		fmt.Println(err.Error())
	}
	fmt.Println(string(output))
	// }
	//
	// err = db.AutoMigrate(&model.TaskData{})
	// if err != nil {
	// 	panic("erro migrating")
	// }
	//
	// router := gin.Default()
	// routes.InitGroup(&router.RouterGroup)
	// err = router.Run(":8080")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
