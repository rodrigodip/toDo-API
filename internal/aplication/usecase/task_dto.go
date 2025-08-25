package usecase

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskDtoOutput struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"decription"`
	Completed   bool   `json:"completed"`
}
