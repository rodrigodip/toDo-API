package main

import (
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/infra/db/mysql"
	"github.com/rodrigodip/toDo-API/internal/infra/repository"
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
	db, err := mysql.NewDataBaseConnection()
	if err != nil {
		fmt.Println(err.Error())
	}
	repo := repository.NewTaskRepositoryDB(db)
	service := usecase.Newtask(repo)
	imput := usecase.CreateTaskRequest{
		Title:       "Minha primeira task",
		Description: "ela funciona!",
	}
	output, err := service.Create(imput)
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonOutput, err := json.Marshal(output)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Saved: %s", jsonOutput)
	//
	// db, err := database.GetDB()
	// if err != nil {
	// 	panic("error getting connection")
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
