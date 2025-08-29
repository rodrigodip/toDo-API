package dependencies

import (
	"fmt"
	"log"
	"os"

	"github.com/rodrigodip/toDo-API/internal/api/http/controller"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/infra/repository"
	"gorm.io/gorm"
)

// TaskRepositoryFactory creates a TaskRepository from config/env
func TaskRepositoryFactory(database *gorm.DB) (repository.TaskRepository, error) {
	persistenceType := os.Getenv("PERSISTENCE_TYPE")
	persistenceFilePath := os.Getenv("FILE_PATH")

	switch persistenceType {
	case "file":
		if persistenceFilePath == "" {
			return nil, fmt.Errorf("FILE_PATH not set for file persistence")
		}
		log.Println("Persistence Method: File System")
		return repository.NewTaskReposytoryFS(persistenceFilePath), nil

	case "mysql":
		if database == nil {
			return nil, fmt.Errorf("database connection is nil for mysql persistence")
		}
		log.Println("Persistence Method: MySQL")
		return repository.NewTaskRepositoryDB(database), nil

	default:
		return nil, fmt.Errorf("invalid persistence type: %s", persistenceType)
	}
}

// Init wires dependencies and returns a TaskController
func Init(database *gorm.DB) (controller.TaskController, error) {
	repo, _ := TaskRepositoryFactory(database)
	// if err != nil {
	// 	return controller.TaskController{}, err
	// }

	service := usecase.NewTaskRepository(repo)
	return controller.NewTaskController(*service), nil
}
