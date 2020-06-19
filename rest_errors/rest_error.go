package rest_errors

import (
	"fmt"
	"net/http"
)

type RestError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restError struct {
	message string        `json:"message"`
	status  int           `json:"code"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e restError) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [%v]",
		e.message, e.status, e.error, e.causes)
}

func (e restError) Message() string {
	return e.message
}
func (e restError) Status() int {
	return e.status
}

func (e restError) Causes() []interface{} {
	return e.causes
}

func NewRestError(msg string, status int, err string, causes []interface{}) RestError {
	return restError{
		message: msg,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

func NewBadRequestError(message string) RestError {
	return restError{
		message: message,
		status:  http.StatusBadRequest,
		error:   http.StatusText(http.StatusBadRequest),
	}
}

func NewNotFoundError(message string) RestError {
	return restError{
		message: message,
		status:  http.StatusNotFound,
		error:   http.StatusText(http.StatusNotFound),
	}
}

func NewInternalServerError(message string, err error) RestError {
	result := restError{
		message: message,
		status:  http.StatusInternalServerError,
		error:   http.StatusText(http.StatusInternalServerError),
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
	}
	return result
}
