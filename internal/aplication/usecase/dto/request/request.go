package request

type TaskRequest struct {
	// @json
	// @jsonTag title
	// @jsonExample Mow the lawn
	// @binding required,min=3,max=30
	Title string `json:"title" binding:"required,min=3,max=30"`
	// @json
	// @jsonTag description
	// @jsonExample Buy fuel for the lawn mower
	// @binding min=3,max=30
	Description string `json:"description,omitempty" binding:"max=50"`
}
