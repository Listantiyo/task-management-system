package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"task-management-system/internal/delivery/dto"
)

// Entity
type User struct {
	ID           uuid.UUID
	Email        string
	Username     string
	PasswordHash string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (u *User) GetID() string       { return u.ID.String() }
func (u *User) GetEmail() string    { return u.Email }
func (u *User) GetUsername() string { return u.Username }
func (u *User) SetPassword(password string) error {

	bytePass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashPass := string(bytePass)
	u.PasswordHash = hashPass
	return nil
}
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// Usecase Interface
type UserUsecase interface {
	Login(ctx context.Context, reqLogin dto.LoginRequest) (*dto.LoginResponse, error)
	Register(ctx context.Context, reqRegister dto.RegisterRequest) (*dto.RegisterResponse, error)
	// UpdateUser(ctx context.Context, reqUpdate dto.UpdateUserRequest) error
}

// Repo Interface
type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
}
