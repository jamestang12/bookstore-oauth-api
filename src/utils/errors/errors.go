package errors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Message: message,
		Error:   "bad request",
	}
}

func NewFoundError(msg string) error {
	return errors.New(msg)
}

func NewBadNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Message: message,
		Error:   "not found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusInternalServerError,
		Message: message,
		Error:   "Internal server error",
	}
}
