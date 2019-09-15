package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

type ErrorResponse struct {
	Error string `json:"error"`
}
type ErrorResponseDev struct {
	ErrorResponse
	Execption error `json:"exception"`
}

func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	response, _ := json.Marshal(data)
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func RespondErrorJSON(w http.ResponseWriter, status int, msg string, exception error) {
	message := ErrorResponse{Error: msg}
	response := ErrorResponseDev{
		ErrorResponse: message,
	}
	if os.Getenv("ENV") == "dev" {
		response.Execption = exception
	}

	jsonResponse, _ := json.Marshal(response)

	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonResponse)
}
