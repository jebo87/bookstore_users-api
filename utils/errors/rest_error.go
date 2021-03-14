package errors

import (
	"fmt"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string, params ...interface{}) *RestErr {
	return &RestErr{
		Message: fmt.Sprintf(message, params...),
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string, params ...interface{}) *RestErr {
	return &RestErr{
		Message: fmt.Sprintf(message, params...),
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewServerError(message string, params ...interface{}) *RestErr {
	return &RestErr{
		Message: fmt.Sprintf(message, params...),
		Status:  http.StatusNotFound,
		Error:   "internal_server_error",
	}
}
