package pkg

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Error struct {
	StatusCode int
	ErrCode    string
	Message    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v-%v-%v", e.StatusCode, e.ErrCode, e.Message)
}

func NewErrBadRequest(msg string) *Error {
	return &Error{
		StatusCode: http.StatusBadRequest,
		ErrCode:    "ERR_BAD_REQUEST",
		Message:    msg,
	}
}

func NewErrNotFound() *Error {
	return &Error{
		StatusCode: http.StatusNotFound,
		ErrCode:    "ERR_NOT_FOUND",
		Message:    "not found",
	}
}

func NewErrInternalError(err error) *Error {
	return &Error{
		StatusCode: http.StatusInternalServerError,
		ErrCode:    "ERR_INTERNAL_ERROR",
		Message:    err.Error(),
	}
}

func NewErrMethodNotAllowed() *Error {
	return &Error{
		StatusCode: http.StatusMethodNotAllowed,
		ErrCode:    "ERR_METHOD_NOT_ALLOWED",
		Message:    "method is not allowed",
	}
}

func NewErrorResp(err error) ApiResp {
	var e *Error
	if !errors.As(err, &e) {
		e = NewErrInternalError(err)
	}
	return ApiResp{
		StatusCode: e.StatusCode,
		OK:         false,
		ErrCode:    e.ErrCode,
		Message:    e.Message,
		Timestamp:  time.Now().Unix(),
	}
}
