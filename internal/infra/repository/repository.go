package repository

import (
	"slices"

	"github.com/rodrigodip/toDo-API/src/model"
)

var (
	TaskRepository = []model.TaskData{}
)

func DeleteTask(index int) {
	TaskRepository = slices.Delete(TaskRepository, index, (index + 1))
}
