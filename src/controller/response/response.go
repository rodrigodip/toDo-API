package response

type TaskResponse struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Added       *string `json:"added"`
	Completed   bool    `json:"completed"`
}
