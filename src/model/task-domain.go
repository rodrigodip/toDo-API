package model

import "gorm.io/gorm"

type TaskData struct {
	gorm.Model
	ID          uint `gorm:"primarykey"`
	Title       string
	Description string
	Completed   bool `gorm:"default=false"`
}

func (td *TaskData) SetId(id uint) {

	td.ID = id
}

func (td *TaskData) GetId() uint {
	return td.ID
}

func (td *TaskData) SetTitle(title string) {
	td.Title = title
}

func (td *TaskData) SetDescription(description string) {
	td.Description = description
}

func (td *TaskData) SetCompleted(completed bool) {
	td.Completed = completed
}

func (td *TaskData) GetAll() TaskData {
	return TaskData{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Completed:   td.Completed,
	}
}
