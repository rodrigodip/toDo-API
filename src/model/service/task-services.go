package service

import (
	"fmt"
	"slices"
	"strconv"
	"sync"

	"github.com/rodrigodip/toDo-API/src/config/rest-err"
	"github.com/rodrigodip/toDo-API/src/controller/model/request"
	"github.com/rodrigodip/toDo-API/src/controller/model/response"
	"github.com/rodrigodip/toDo-API/src/model"
)

var (
	tasks   = []model.TaskData{}
	nextID  = 1
	taskMux sync.Mutex
)

func CreateTask(rep request.TaskRequest) response.TaskRespose {
	taskMux.Lock()
	defer taskMux.Unlock()

	var newTask model.TaskData

	newTask.SetId(nextID)
	newTask.SetTitle(rep.Title)
	newTask.SetDescription(rep.Description)

	nextID++
	tasks = append(tasks, newTask)

	return response.TaskRespose{
		ID:          newTask.ID,
		Title:       newTask.Title,
		Description: newTask.Description,
		Completed:   newTask.Completed,
	}
}

func GetAllTasks() []model.TaskData {
	return tasks
}

func GetTaskByID(id string) (model.TaskData, *rest_err.RestErr) {

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return model.TaskData{}, restErr
	}

	for _, t := range tasks {
		if t.ID == intId {
			return t, nil
		}
	}

	restErr := rest_err.NewNotFoundError(
		fmt.Sprintln("Error: ID not Found"),
	)
	return model.TaskData{}, restErr
}

func UpdateTaskByID(id string, updated model.TaskData) (model.TaskData, *rest_err.RestErr) {
	taskMux.Lock()
	defer taskMux.Unlock()

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return model.TaskData{}, restErr
	}

	for i, t := range tasks {
		if t.ID == intId {
			tasks[i].Title = updated.Title
			tasks[i].Description = updated.Description
			tasks[i].Completed = updated.Completed
			return tasks[i], nil
		}
	}
	restErr := rest_err.NewNotFoundError(
		fmt.Sprintln("Error: ID not Found"),
	)
	return model.TaskData{}, restErr
}

func DeleteTaskByID(id string) *rest_err.RestErr {
	taskMux.Lock()
	defer taskMux.Unlock()

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return restErr
	}

	for i, t := range tasks {
		if t.ID == intId {
			tasks = slices.Delete(tasks, i, (i + 1))
			return nil
		}
	}
	restErr := rest_err.NewNotFoundError(
		fmt.Sprintln("Error: Task not Found"),
	)
	return restErr
}

func SetTaskDone(id string) *rest_err.RestErr {
	taskMux.Lock()
	defer taskMux.Unlock()

	intId, err := strconv.Atoi(id) // converts string to integer
	if err != nil {
		restErr := rest_err.NewBadRequest(
			fmt.Sprintln("Error: ID must be a number"),
		)
		return restErr
	}

	for i, t := range tasks {
		if t.ID == intId {
			tasks[i].Completed = true
			return nil
		}
	}
	restErr := rest_err.NewNotFoundError(
		fmt.Sprintln("Error: Task not Found"),
	)
	return restErr
}
