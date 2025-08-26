package repository

import (
	"gorm.io/gorm"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}
type taskRepositoryDB struct {
	mysqlDB *gorm.DB
}

func NewTaskRepositoryDB(database *gorm.DB) *taskRepositoryDB {
	return &taskRepositoryDB{mysqlDB: database}
}

type TaskRepository interface {
	Create(id, title, description string, completed bool) error
}

func (t *taskRepositoryDB) Create(id, title, description string, completed bool) error {

	newTask := Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
	db := t.mysqlDB
	insert := db.Create(&newTask)
	if err := insert.Error; err != nil {
		return err
	}

	return nil
}
