package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	apperr "task-management-system/internal/delivery/error"
	"task-management-system/internal/domain"
	"task-management-system/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var userModel models.UserModel
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperr.New(apperr.ErrRecordNotFound, fmt.Sprintf("user with email `%s` not found", email))
		}
		slog.Error("error get account by email", "error", err)
		return nil, apperr.New(apperr.ErrInternalServer, "failed to get user")
	}

	user := &domain.User{
		ID:           userModel.ID,
		Email:        userModel.Email,
		Username:     userModel.Username,
		PasswordHash: userModel.PasswordHash,
		CreatedAt:    &userModel.CreatedAt,
		UpdatedAt:    &userModel.UpdatedAt,
	}
	return user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userModel := models.UserModel{
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}
	err := r.db.WithContext(ctx).Create(&userModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, apperr.New(apperr.ErrEmailAlreadyExits, "invalid duplicate email")
		}
		slog.Error("error create user", "error", err)
		return nil, err
	}

	createdUser := &domain.User{
		ID:           userModel.ID,
		Email:        userModel.Email,
		Username:     userModel.Username,
		PasswordHash: userModel.PasswordHash,
		CreatedAt:    &userModel.CreatedAt,
	}

	return createdUser, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	updateData := map[string]any{
		"email":    user.Email,
		"username": user.Username,
	}
	if err := r.db.WithContext(ctx).Model(&models.UserModel{ID: user.ID}).Updates(&updateData).Error; err != nil {
		return err
	}
	return nil
}
