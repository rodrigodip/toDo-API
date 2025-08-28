package repository

import (
	"github.com/rodrigodip/toDo-API/internal/domain"
	"gorm.io/gorm"
)

type taskRepositoryDB struct {
	mysqlDB *gorm.DB
}

func NewTaskRepositoryDB(database *gorm.DB) *taskRepositoryDB {
	return &taskRepositoryDB{mysqlDB: database}
}

func (t *taskRepositoryDB) Create(id, title, description string, completed bool) error {

	newTask := task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	created := t.mysqlDB.Create(&newTask)
	return created.Error
}

func (t *taskRepositoryDB) GetTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	result := t.mysqlDB.Find(&tasks)
	return tasks, result.Error
}
