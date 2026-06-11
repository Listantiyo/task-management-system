package dto

import "errors"

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
	ErrRecordNotFound ErrorCode = "RECORD_NOT_FOUND"

	// --- MODUL JWT (TAMBAHAN BARU) ---
	ErrTokenMissing ErrorCode = "TOKEN_MISSING" // Jika header Authorization kosong
	ErrTokenInvalid ErrorCode = "TOKEN_INVALID" // Jika format token rusak / signature salah
	ErrTokenExpired ErrorCode = "TOKEN_EXPIRED" // Jika token sudah melewati masa aktif
)

// String() membuat Enum ini bisa mencetak string biasa jika dibutuhkan
func (e ErrorCode) String() string {
	return string(e)
}
func (e ErrorCode) Error() string {
	return string(e)
}

// Type Assertion Error
func AsType[T error](err error) (T, bool) {
	var target T
	if errors.As(err, &target) {
		return target, true
	}
	return target, false
}

// Response godoc
// @name Response
type Response[T any] struct {
	Success bool       `json:"success"`
	Data    T          `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
	Meta    *Meta      `json:"meta,omitempty"`
}

// ErrorInfo godoc
// @name ErrorInfo
type ErrorInfo struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// Meta godoc
// @name Meta
type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}
