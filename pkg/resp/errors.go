package resp

import "net/http"

type AppError struct {
	Status  int
	Code    string
	Message string
	Data    any
}

func (e *AppError) Error() string {
	return e.Message
}

func E(status int, code, message string, data any) *AppError {
	return &AppError{Status: status, Code: code, Message: message, Data: data}
}

func ErrBadRequest(code, msg string, data any) *AppError {
	return E(http.StatusBadRequest, code, msg, data)
}

func ErrUnauthorized(code, msg string) *AppError {
	return E(http.StatusUnauthorized, code, msg, nil)
}

func ErrForbidden(code, msg string) *AppError {
	return E(http.StatusForbidden, code, msg, nil)
}

func ErrNotFound(code, msg string) *AppError {
	return E(http.StatusNotFound, code, msg, nil)
}

func ErrTooMany(code, msg string) *AppError {
	return E(http.StatusTooManyRequests, code, msg, nil)
}

func ErrInternal(code, msg string) *AppError {
	return E(http.StatusInternalServerError, code, msg, nil)
}
