package dependencies

import (
	"github.com/rodrigodip/toDo-API/internal/api/http/controller"
	"github.com/rodrigodip/toDo-API/internal/aplication/usecase"
	"github.com/rodrigodip/toDo-API/internal/infra/repository"
	"gorm.io/gorm"
)

func Init(
	database *gorm.DB,
) controller.TaskController {
	repo := repository.NewTaskRepositoryDB(database)
	service := usecase.Newtask(repo)
	return controller.NewTaskController(*service)
}
