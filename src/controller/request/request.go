package request

type TaskRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=20"`
	Description string `json:"description,omitempty"`
}
