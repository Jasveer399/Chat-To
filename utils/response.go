package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Status    int         `json:"status"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
	Timestamp string      `json:"timestamp"`
}

type ErrorResponse struct {
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Error     string      `json:"error"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp string      `json:"timestamp"`
}

func SendResponse(w http.ResponseWriter, status int, message string, data interface{}, meta interface{}) {
	res := Response{
		Status:    status,
		Success:   true,
		Message:   message,
		Data:      data,
		Meta:      meta,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func SendError(w http.ResponseWriter, status int, message string, err error, data interface{}) {
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	} else {
		errMsg = message
	}
	res := ErrorResponse{
		Status:    status,
		Message:   message,
		Error:     errMsg,
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
