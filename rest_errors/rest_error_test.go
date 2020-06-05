package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("my message", errors.New("database error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "my message", err.Message)
	assert.EqualValues(t, "Internal Server Error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "database error", err.Causes[0])
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("my message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "my message", err.Message)
	assert.EqualValues(t, "Not Found", err.Error)
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("my message")

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "my message", err.Message)
	assert.EqualValues(t, "Bad Request", err.Error)
}
