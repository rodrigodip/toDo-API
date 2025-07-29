package model

import (
	"fmt"
	"time"
)

type TaskDomain struct {
	id          int
	title       string
	description string
	added       time.Time
	completed   bool
}

func (td *TaskDomain) GetID() int {
	return td.id
}

func (td *TaskDomain) SetID(id int) {
	td.id = id
}

func (td *TaskDomain) GetTitle() string {
	return td.title
}

func (td *TaskDomain) GetDescription() string {
	return td.description
}
func (td *TaskDomain) GetAdded() time.Time {
	return td.added
}

func (td *TaskDomain) SetAdded() {
	t := time.Now()
	td.added = t
	fmt.Println(t)
}
