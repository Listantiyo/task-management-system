package dto

import apperr "task-management-system/internal/delivery/error"

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
	Code    apperr.ErrorCode `json:"code"`
	Message string           `json:"message"`
}

// Meta godoc
// @name Meta
type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}
