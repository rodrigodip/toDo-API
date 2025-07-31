package request

type TaskRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=30"`
	Description string `json:"description,omitempty" binding:"max=50"`
}
