package repository

import (
	"github.com/rodrigodip/toDo-API/src/model"
)

var TaskRepo map[int]model.TaskDomain

func init() {
	TaskRepo = make(map[int]model.TaskDomain)
}
