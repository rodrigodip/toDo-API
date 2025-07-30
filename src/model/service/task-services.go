package service

import (
	"errors"
	"github.com/rodrigodip/toDo-API/src/model"
	"slices"
	"sync"
)

var (
	tasks   = []model.Task{}
	nextID  = 1
	taskMux sync.Mutex
)

func CreateTask(task model.Task) model.Task {
	taskMux.Lock()
	defer taskMux.Unlock()

	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return task
}

func GetAllTasks() []model.Task {
	return tasks
}

func GetTaskByID(id int) (model.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func UpdateTaskByID(id int, updated model.Task) (model.Task, error) {
	taskMux.Lock()
	defer taskMux.Unlock()

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Title = updated.Title
			tasks[i].Description = updated.Description
			tasks[i].Completed = updated.Completed
			return tasks[i], nil
		}
	}
	return model.Task{}, errors.New("task not found")
}

func DeleteTaskByID(id int) error {
	taskMux.Lock()
	defer taskMux.Unlock()

	for i, t := range tasks {
		if t.ID == id {
			tasks = slices.Delete(tasks, i, (i + 1))
			return nil
		}
	}
	return errors.New("task not found")
}
