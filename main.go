package stdmodels

import (
	"encoding/json"
	"fmt"
	"gitlab.com/developerdurp/logger"
	"net/http"
)

type BasicMessage struct {
	Message string `json:"message"`
}

type StandardMessage struct {
	Message interface{}
	Status  int `json:"status"`
}

type StandardError struct {
	Message     string   `json:"message"`
	Status      int      `json:"status"`
	Description []string `json:"description"`
}

// NewFailureResponse returns a new instance of StandardError with the given message, status code and description.
func NewFailureResponse(message string, status int, description []string) StandardError {
	return StandardError{
		Message:     message,
		Status:      status,
		Description: description,
	}
}

// NewMessageResponse returns a new instance of StandardMessage with the given message and status code.
func NewMessageResponse(message interface{}, status int) StandardMessage {
	return StandardMessage{
		Message: message,
		Status:  status,
	}
}

func SendResponse(message interface{}, w http.ResponseWriter) {
	switch req := message.(type) {
	case *StandardMessage:

		// Write the status code to the response.
		w.WriteHeader(req.Status)
		w.Header().Set("Content-Type", "application/json")

		// Write the message to the response body.
		err := json.NewEncoder(w).Encode(req.Message)
		if err != nil {
			logger.LogError("Failed to Encode")
		}
	case *StandardError:

		// Write the status code to the response.
		w.WriteHeader(req.Status)
		w.Header().Set("Content-Type", "application/json")

		// Write the message to the response body.
		err := json.NewEncoder(w).Encode(req)
		if err != nil {
			logger.LogError("Failed to Encode")
		}
	default:
		fmt.Println("Passed wrong interface")
	}
}
