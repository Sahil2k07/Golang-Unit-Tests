package errors

import "net/http"

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

func (e *CustomError) Error() string {
	return e.Message
}

func BadRequestError(message string) *CustomError {
	return &CustomError{
		Success:    false,
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func NotFoundError(message string) *CustomError {
	return &CustomError{
		Success:    false,
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func InternalServerError(message string) *CustomError {
	return &CustomError{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func UnAuthorizedError(message string) *CustomError {
	return &CustomError{
		Success:    false,
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}
}
