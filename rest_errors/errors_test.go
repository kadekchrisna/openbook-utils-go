package errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalError(t *testing.T) {
	err := NewInternalServerError("this is message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, errors.New("database error"), err.Causes[0])

	err = NewInternalServerError("this is message", nil)
	assert.EqualValues(t, errors.New("something when wrong i guess"), err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {
	// TODO
	err := NewNotFoundError("this is message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)
}

func TestNewNotFoundError(t *testing.T) {
	// TODO
	err := NewBadRequestError("this is message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)
}

func TestNewError(t *testing.T) {
	// TODO
}
