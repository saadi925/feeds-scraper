package utils

import (
	"encoding/json"
	"net/http"
)
// Sets the Content Type to application/json by default , takes response writer , data interface{}, and status code int. 
// like  w,data,500
func RespondWithJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Responds with a message and a status code   
// type ErrorResponse struct {
// 	Message string `json:"message"`
// 	Code    int    `json:"code"`
// }
func RespondWithError(w http.ResponseWriter, err error, status int) {
	errorResponse := NewAppError(status, err.Error())
	HandleError(w, errorResponse)
}

// ParseJSON parses JSON data from a request and binds it to a target struct.
func ParseJSON(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(target)
}