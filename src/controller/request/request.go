package request

type TaskResponse struct {
	TaskId      uint8
	Title       string
	Description *string
	Added       *string
	Completed   bool
}
