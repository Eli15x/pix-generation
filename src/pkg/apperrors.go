package apperrors

import (
	"errors"
	"net/http"
)

type Code string

const (
	CodeNotFound     Code = "not_found"
	CodeConflict     Code = "conflict"
	CodeValidation   Code = "validation"
	CodeUnauthorized Code = "unauthorized"
	CodeInternal     Code = "internal"
)

type AppError struct {
	Code Code   `json:"code"`
	Msg  string `json:"message,omitempty"`
	Err  error  `json:"-"`
}

type PublicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Public() PublicError {
	if e == nil {
		return PublicError{
			Code:    string(CodeInternal),
			Message: "internal server error",
		}
	}
	msg := e.Msg
	if msg == "" {
		switch e.Code {
		case CodeNotFound:
			msg = "resource not found"
		case CodeValidation:
			msg = "validation error"
		case CodeUnauthorized:
			msg = "unauthorized"
		case CodeConflict:
			msg = "conflict"
		default:
			msg = "internal server error"
		}
	}
	return PublicError{
		Code:    string(e.Code),
		Message: msg,
	}
}

func (e *AppError) Error() string {
	if e == nil {
		return "<nil>"
	}
	if e.Err != nil && e.Msg != "" {
		return e.Msg + ": " + e.Err.Error()
	}
	if e.Msg != "" {
		return e.Msg
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Code)
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.Code == t.Code
}

func New(code Code, msg string) *AppError {
	return &AppError{Code: code, Msg: msg}
}

func Wrap(err error, code Code, msg string) *AppError {
	if err == nil {
		return New(code, msg)
	}
	return &AppError{Code: code, Msg: msg, Err: err}
}

var (
	ErrNotFound     = &AppError{Code: CodeNotFound}
	ErrConflict     = &AppError{Code: CodeConflict}
	ErrValidation   = &AppError{Code: CodeValidation}
	ErrUnauthorized = &AppError{Code: CodeUnauthorized}
	ErrInternal     = &AppError{Code: CodeInternal}
)

func HTTPStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	var ae *AppError
	if errors.As(err, &ae) {
		switch ae.Code {
		case CodeNotFound:
			return http.StatusNotFound
		case CodeConflict:
			return http.StatusConflict
		case CodeValidation:
			return http.StatusBadRequest
		case CodeUnauthorized:
			return http.StatusUnauthorized
		case CodeInternal:
			return http.StatusInternalServerError
		default:
			return http.StatusInternalServerError
		}
	}

	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
