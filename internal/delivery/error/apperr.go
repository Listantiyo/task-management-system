package apperr

import (
	"errors"
	"fmt"
)

type ErrorCode string

// Daftar Enum Error Code - Kelompokkan berdasarkan modul agar rapi
const (
	// Modul Autentikasi & Umum
	ErrInvalidInput   ErrorCode = "INVALID_INPUT"
	ErrUnauthorized   ErrorCode = "UNAUTHORIZED"
	ErrInternalServer ErrorCode = "INTERNAL_SERVER_ERROR"

	// Modul User
	ErrUserNotFound      ErrorCode = "USER_NOT_FOUND"
	ErrEmailAlreadyExits ErrorCode = "EMAIL_ALREADY_EXISTS"
	ErrPasswordWrong     ErrorCode = "PASSWORD_INCORRECT"

	// Modul Umum
	ErrRecordNotFound   ErrorCode = "RECORD_NOT_FOUND"
	ErrNoRecordAffected ErrorCode = "NO_RECORD_AFFECTED"

	// --- MODUL JWT (TAMBAHAN BARU) ---
	ErrTokenMissing ErrorCode = "TOKEN_MISSING" // Jika header Authorization kosong
	ErrTokenInvalid ErrorCode = "TOKEN_INVALID" // Jika format token rusak / signature salah
	ErrTokenExpired ErrorCode = "TOKEN_EXPIRED" // Jika token sudah melewati masa aktif
)

type AppError struct {
	Code    ErrorCode
	Message string
	RawErr  error
}

func (e AppError) Error() string {
	if e.RawErr != nil {
		return fmt.Sprintf("[%s] %s | raw: %v", e.Code, e.Message, e.RawErr)
	}

	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func New(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func Wrap(code ErrorCode, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		RawErr:  err,
	}
}

// Type Assertion Error
func AsType[T error](err error) (T, bool) {
	var target T
	if errors.As(err, &target) {
		return target, true
	}
	return target, false
}
