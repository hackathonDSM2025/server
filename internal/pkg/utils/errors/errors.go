package errors

import (
	"encoding/json"
	"net/http"
)

type CustomError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Details interface{} `json:"details,omitempty"`
	}{
		Success: false,
		Message: e.Message,
		Details: e.Details,
	})
}

func NewCustomError(status int, message string, details interface{}) *CustomError {
	return &CustomError{
		Status:  status,
		Message: message,
		Details: details,
	}
}

func BadRequest(message string) *CustomError {
	return NewCustomError(http.StatusBadRequest, message, nil)
}

func Unauthorized(message string) *CustomError {
	return NewCustomError(http.StatusUnauthorized, message, nil)
}

func Forbidden(message string) *CustomError {
	return NewCustomError(http.StatusForbidden, message, nil)
}

func NotFound(message string) *CustomError {
	return NewCustomError(http.StatusNotFound, message, nil)
}

func Conflict(message string) *CustomError {
	return NewCustomError(http.StatusConflict, message, nil)
}

func InternalServerError(message string) *CustomError {
	return NewCustomError(http.StatusInternalServerError, message, nil)
}