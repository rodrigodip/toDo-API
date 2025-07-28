package request

type TaskRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
}
