package model

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required,min=3,max=30"`
	Description string `json:"description,omitempty" binding:"max=50"`
	Completed   bool   `json:"completed"`
}
