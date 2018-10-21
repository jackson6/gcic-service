package handler

import (
	"net/http"
	"encoding/json"
	"log"
)

type HttpResponse struct {
	ResultCode  int `json:"result,omitempty"`
	CodeContent string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

// respondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, status int, err jsonErr, message error) {
	log.Println(message)
	RespondJSON(w, status, err)
}

type jsonErr struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

var InternalError = jsonErr{
	Code: 5000,
	Message: "Internal Server Error",
}

var BadRequest = jsonErr{
	Code: 400,
	Message: "Bad request",
}

var UnauthorizedError = jsonErr{
	Code: 403,
	Message: "Unauthorized request",
}

var NotFound = jsonErr{
	Code: 400,
	Message: "Not Found",
}