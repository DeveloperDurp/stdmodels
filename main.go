package stdmodels

import (
	"encoding/json"
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

type Response interface {
	SendResponse(w http.ResponseWriter)
}

func (message *StandardMessage) SendReponse(w http.ResponseWriter) {
	setHeader(&w, message.Status)

	// Write the message to the response body.
	err := json.NewEncoder(w).Encode(message.Message)
	if err != nil {
		logger.LogError("Failed to Encode")
	}
}

func (message *StandardError) SendReponse(w http.ResponseWriter) {
	setHeader(&w, message.Status)

	// Write the message to the response body.
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		logger.LogError("Failed to Encode")
	}
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

// SetHeader sets the HTTP response headers for a JSON response.
func setHeader(w *http.ResponseWriter, statusCode int) {
	(*w).WriteHeader(statusCode)
	(*w).Header().Set("Content-Type", "application/json")
}
