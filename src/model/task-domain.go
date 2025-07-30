package model

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required,min=3,max=20"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
