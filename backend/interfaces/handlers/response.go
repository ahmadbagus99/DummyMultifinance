package interfaces

import (
	"encoding/json"
	"net/http"
)

var (
	Success     = "00"
	DataFound   = "01"
	DataUpdated = "02"

	ClientError  = "10"
	DataNotFound = "11"
	InvalidToken = "12"
	Unauthorized = "13"

	Timeout     = "20"
	ServerError = "21"
	SystemError = "22"
)

type Response struct {
	ResultCode string      `json:"resultCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func GeneralResponse(w http.ResponseWriter, statusCode int, resultCode, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := Response{
		ResultCode: resultCode,
		Message:    message,
		Data:       data,
	}

	json.NewEncoder(w).Encode(response)
}
