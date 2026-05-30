package dto

import (
	"time"

	"github.com/google/uuid"
)

// DTO Request

// RegisterRequest godoc
// @name RegisterRequest
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=2"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest godoc
// @name LoginRequest
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UpdateUserRequest godoc
// @name UpdateUserRequest
type UpdateUserRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Username        string `json:"username" binding:"required,min=2"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=6"`
}

// DTO Response

// RegisterResponse godoc
// @name RegisterResponse
type RegisterResponse struct {
	ID        uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
}

// LoginResponse godoc
// @name LoginResponse
type LoginResponse struct {
	ID       uuid.UUID `json:"user_id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
}
