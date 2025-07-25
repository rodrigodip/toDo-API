package config

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
}

func NewRestError(message, err string, code int) *RestError {
	return &RestError{
		Message: message,
		Err:     err,
		Code:    code,
	}
}
func NewBadRequest(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "Bad Request",
		Code:    http.StatusBadRequest,
	}
}
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "Task not Found",
		Code:    http.StatusNotFound,
	}
}
