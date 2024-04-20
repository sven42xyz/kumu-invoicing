package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
}

func WriteResponse(w http.ResponseWriter, message string, statusCode int) {
	r := Response{
		Message:    message,
		StatusCode: statusCode,
	}

	// Marshal ErrorResponse into JSON
	jr, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response with status code
	w.WriteHeader(statusCode)
	w.Write(jr)
}
