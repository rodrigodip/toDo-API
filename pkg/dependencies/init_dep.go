package dependencies

import (
	"github.com/rodrigodip/toDo-API/internal/api/http/controller"
	"github.com/rodrigodip/toDo-API/internal/api/http/handler"
	"github.com/rodrigodip/toDo-API/internal/infra/repository"
	"gorm.io/gorm"
)

func Init(
	database *gorm.DB,
) controller.TaskController {
	repo := repository.NewTaskRepositoryDB(database)
	handler := handler.NewTaskHandler(repo)
	return controller.NewTaskController(handler)
}
