package domain

import (
	"errors"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) TaskValidation() error {
	if t.Title == "" {
		return errors.New("O título é obnrigatório")
	}
	if len(t.Title) > 30 {
		return errors.New("O titulo não pode ter mais que 30 characters.")
	}
	// NOTE: Description é opcional, preciso checar existência primeiro?
	if len(t.Description) > 50 {
		return errors.New("A Description não pode ter mais que 50 characters.")
	}
	return nil
}
