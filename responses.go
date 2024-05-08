package stdmodels

import (
	"encoding/json"
	"net/http"

	"gitlab.com/developerdurp/logger"
)

func FailureReponse(message string, w http.ResponseWriter, statusCode int, description []string) {
	response := StandardError{
		Message:     message,
		Status:      statusCode,
		Description: description,
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.LogError("Failed to Encode")
	}
}

func SuccessResponse(message string, w http.ResponseWriter, statusCode int) {
	response := StandardMessage{
		Message: message,
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.LogError("Failed to Encode")
	}
}
