package dependencies

import (
	"log"
	"os"

	"github.com/rodrigodip/toDo-API/internal/api/http/controller"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/infra/repository"
	"gorm.io/gorm"
)

func Init(
	database *gorm.DB,
) controller.TaskController {

	persistenceType := os.Getenv("PERSISTENCE_TYPE")
	persistenceFilePath := os.Getenv("FILE_PATH")

	var repo repository.TaskRepository
	switch persistenceType {
	case "file":
		repo = repository.NewTaskReposytoryFS(persistenceFilePath)
		log.Println("Persistence Method: File System")
	case "mysql":
		repo = repository.NewTaskRepositoryDB(database)
		log.Println("Persistence Method: MySQL")
	default:
		log.Fatal("Método de persistência não definido")
	}

	//repo := repository.NewTaskRepositoryDB(database)
	service := usecase.Newtask(repo)
	return controller.NewTaskController(*service)
}
