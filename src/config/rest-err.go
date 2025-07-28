package rest_err

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

func (r *RestErr) Error() string {
	return r.Message
}
func NewRestError(message, err string, code int) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
	}
}
func NewBadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Task not Found",
		Code:    http.StatusNotFound,
	}
}
