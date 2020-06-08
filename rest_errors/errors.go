package errors

import (
	"errors"
	"net/http"
)

type ResErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

// NewBadRequestError for error template with status 400
func NewBadRequestError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}

}

func NewRestError(message string, status int, err string, causes []interface{}) *ResErr {
	return &ResErr{
		Message: message,
		Status:  status,
		Error:   "bad_request",
		Causes:  causes,
	}
}

func NewUnAuthorizedError() *ResErr {
	return &ResErr{
		Message: "access token is not valid",
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}

// NewNotFoundError for error template with status 40
func NewNotFoundError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError something went wrong
func NewInternalServerError(message string, err error) *ResErr {
	if err == nil {
		err = errors.New("something when wrong i guess")
	}
	return &ResErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		Causes:  []interface{}{err},
	}
}
