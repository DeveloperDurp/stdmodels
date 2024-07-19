package stdmodels

import (
	"encoding/json"
	"gitlab.com/developerdurp/logger"
	"net/http"
)

type StandardMessage struct {
	Message string `json:"message"`
}

type StandardError struct {
	Message     string   `json:"message"`
	Status      int      `json:"status"`
	Description []string `json:"description"`
}

func NewFailureResponse(message string, status int, description []string) StandardError {
	return StandardError{
		Message:     message,
		Status:      status,
		Description: description,
	}
}

func NewMessageResponse(message string) StandardMessage {
	return StandardMessage{
		Message: message,
	}
}

func FailureReponse(StandardError StandardError, w http.ResponseWriter) {

	w.WriteHeader(StandardError.Status)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(StandardError)
	if err != nil {
		logger.LogError("Failed to Encode")
	}
}

func SuccessResponse(message interface{}, w http.ResponseWriter, statusCode int) {

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		logger.LogError("Failed to Encode")
	}
}
