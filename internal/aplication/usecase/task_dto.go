package usecase

type CreateTaskRequest struct {
	Title       string
	Description string
}

type TaskDtoOutput struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}
