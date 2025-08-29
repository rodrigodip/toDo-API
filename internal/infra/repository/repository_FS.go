package repository

import (
	"encoding/json"
	"errors"
	"sync"

	"os"

	"github.com/rodrigodip/toDo-API/internal/domain"
)

type taskRepositoryFS struct {
	filePath string
	mutex    sync.Mutex
}

func NewTaskReposytoryFS(filePath string) *taskRepositoryFS {
	return &taskRepositoryFS{filePath: filePath}
}

// Create writes a new task task into tasks.txt
func (r *taskRepositoryFS) Create(id, title, description string, completed bool) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	newTask := task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	file, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(newTask)
}

// GetTasks retrieves all tasks from tasks.txt
func (r *taskRepositoryFS) GetTasks() ([]domain.Task, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Open(r.filePath)
	if err != nil {
		// If file doesn't exist yet, return empty
		if os.IsNotExist(err) {
			return []domain.Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []domain.Task
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var task domain.Task
		if err := decoder.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
func (r *taskRepositoryFS) GetTask(id string) (domain.Task, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	file, err := os.Open(r.filePath)
	if err != nil {
		// If file doesn't exist yet, return empty
		if os.IsNotExist(err) {
			return domain.Task{}, errors.New("NotFoundError: File not found")
		}
		return domain.Task{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var task domain.Task
		if err := decoder.Decode(&task); err != nil {
			return domain.Task{}, err
		}
		if task.ID == id {
			return task, nil
		}
	}
	return domain.Task{}, errors.New("NotFoundError: Invalid ID")

}
func (r *taskRepositoryFS) UpdateTask(id, title, description string) (domain.Task, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Read all tasks
	file, err := os.Open(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return domain.Task{}, errors.New("NotFoundError: File not found")
		}
		return domain.Task{}, err
	}
	defer file.Close()

	var tasks []domain.Task
	decoder := json.NewDecoder(file)
	for decoder.More() {
		var task domain.Task
		if err := decoder.Decode(&task); err != nil {
			return domain.Task{}, err
		}
		tasks = append(tasks, task)
	}

	// Find and update the task
	var updatedTask domain.Task
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			tasks[i].Description = description
			updatedTask = tasks[i]
			found = true
			break
		}
	}

	if !found {
		return domain.Task{}, errors.New("NotFoundError: Invalid ID")
	}

	// Write all tasks back to file
	tempFile := r.filePath + ".tmp"
	newFile, err := os.Create(tempFile)
	if err != nil {
		return domain.Task{}, errors.New("Error: Invalid ID")
	}
	defer newFile.Close()

	encoder := json.NewEncoder(newFile)
	for _, task := range tasks {
		if err := encoder.Encode(task); err != nil {
			return domain.Task{}, err
		}
	}

	// Replace the original file with the updated one
	if err := os.Rename(tempFile, r.filePath); err != nil {
		return domain.Task{}, err
	}

	return updatedTask, nil
}

func (t *taskRepositoryFS) DeleteTask(id string) error {

	return nil
}
func (t *taskRepositoryFS) SetTaskDone(id string) error {

	return nil
}
