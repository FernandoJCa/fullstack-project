package utils

import (
	"log"
	"net/http"
)

type AppError struct {
	Error   error
	Message string
	Code    int
}

func (ae *AppError) LogAndRespond(w http.ResponseWriter) {
	log.Printf("Error: %v", ae.Error)
	http.Error(w, ae.Message, ae.Code)
}

func InternalServerError(err error) *AppError {
	return &AppError{
		Error:   err,
		Message: "Internal Server Error",
		Code:    http.StatusInternalServerError,
	}
}

func BadRequestError(err error) *AppError {
	return &AppError{
		Error:   err,
		Message: "Bad Request",
		Code:    http.StatusBadRequest,
	}
}

func NotFoundError(err error) *AppError {
	return &AppError{
		Error:   err,
		Message: "Not Found",
		Code:    http.StatusNotFound,
	}
}
