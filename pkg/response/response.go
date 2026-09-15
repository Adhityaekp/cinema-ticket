package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func JSON(
	w http.ResponseWriter,
	statusCode int,
	message string,
	data interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(
	w http.ResponseWriter,
	statusCode int,
	message string,
	err interface{},
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

func Success(
	w http.ResponseWriter,
	message string,
	data interface{},
) {
	JSON(w, http.StatusOK, message, data)
}

func Created(
	w http.ResponseWriter,
	message string,
	data interface{},
) {
	JSON(w, http.StatusCreated, message, data)
}
	