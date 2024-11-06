package helpers

import (
	"errors"
	"fmt"
)

var (
	ErrUnexpected error = errors.New("error.unexpected")
)

type CustomError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *CustomError) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%d - %s", e.StatusCode, e.Message)
}

// Funções auxiliares para criar erros específicos
func NewBadRequestError(message string) *CustomError {
	return &CustomError{
		StatusCode: 400,
		Message:    message,
	}
}

func NewUnprocessableEntityError(message string) *CustomError {
	return &CustomError{
		StatusCode: 422,
		Message:    message,
	}
}

func NewUnauthorizedError(message string) *CustomError {
	return &CustomError{
		StatusCode: 401,
		Message:    message,
	}
}

func NewForbiddenError(message string) *CustomError {
	return &CustomError{
		StatusCode: 403,
		Message:    message,
	}
}

func NewNotFoundError(message string) *CustomError {
	return &CustomError{
		StatusCode: 404,
		Message:    message,
	}
}

func NewConflictError(message string) *CustomError {
	return &CustomError{
		StatusCode: 409,
		Message:    message,
	}
}

func NewInternalServerError(message string) *CustomError {
	return &CustomError{
		StatusCode: 500,
		Message:    message,
	}
}

func NewBadGatewayError(message string) *CustomError {
	return &CustomError{
		StatusCode: 502,
		Message:    message,
	}
}

func NewServiceUnavailableError(message string) *CustomError {
	return &CustomError{
		StatusCode: 503,
		Message:    message,
	}
}

// Funções auxiliares para criar respostas de sucesso
func NewOKResponse(message string) *CustomError {
	return &CustomError{
		StatusCode: 200,
		Message:    message,
	}
}

func NewCreatedResponse(message string) *CustomError {
	return &CustomError{
		StatusCode: 201,
		Message:    message,
	}
}

func NewNoContentResponse(message string) *CustomError {
	return &CustomError{
		StatusCode: 204,
		Message:    message,
	}
}
