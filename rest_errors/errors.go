package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type resErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

type ResErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

func (e resErr) Error() string {
	return fmt.Sprintf("message: %s, status: %d, error: %s, causes [ %v ]", e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e resErr) Message() string {
	return e.ErrMessage
}

func (e resErr) Status() int {
	return e.ErrStatus
}

func (e resErr) Causes() []interface{} {
	return e.ErrCauses
}

func NewError(message string) error {
	return errors.New(message)
}

// NewBadRequestError for error template with status 400
func NewBadRequestError(message string) ResErr {
	return resErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}

}

func NewRestError(message string, status int, err string, causes []interface{}) ResErr {
	return resErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   "bad_request",
		ErrCauses:  causes,
	}
}

func NewUnAuthorizedError() ResErr {
	return resErr{
		ErrMessage: "access token is not valid",
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

// NewNotFoundError for error template with status 40
func NewNotFoundError(message string) ResErr {
	return resErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

// NewInternalServerError something went wrong
func NewInternalServerError(message string, err error) ResErr {
	result := resErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
