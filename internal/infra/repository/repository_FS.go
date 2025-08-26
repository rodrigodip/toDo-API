package repository

import (
	"encoding/json"
	"sync"

	"os"
)

type taskRepositoryFS struct {
	filePath string
	mutex    sync.Mutex
}

func NewTaskReposytoryFS(filePath string) *taskRepositoryFS {
	return &taskRepositoryFS{filePath: filePath}
}

// type TaskRepository interface {
// 	Create(id, title, description string, completed bool) error
// }

func (r *taskRepositoryFS) Create(id, title, description string, completed bool) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	newTask := Task{
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
