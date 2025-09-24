package helper

import "net/http"

type AppError struct {
	StatusCode      int
	ValidationError any
	Message         string
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(statusCode int, message string) *AppError {
	return &AppError{
		StatusCode:      statusCode,
		ValidationError: nil,
		Message:         message,
	}
}

func InvalidJsonError() *AppError {
	return BadRequestError("invalid json")
}

func BadRequestError(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NotFoundError(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}
