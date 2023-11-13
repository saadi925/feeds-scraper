// utils/errors.go
package utils

import (
	"encoding/json"
	"net/http"
)

// for Error Response
type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type AppError struct {
	Code    int
	Message string
}
//New App Error Returns a pointer to the App Error
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}
// handles the error response , with message and code
func HandleError(w http.ResponseWriter, err *AppError) {

	errorResponse := ErrorResponse{
		Message: err.Message,
		Code:    err.Code,
	}
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(errorResponse)
}
