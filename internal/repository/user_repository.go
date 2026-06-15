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
	user, err := gorm.G[models.UserModel](r.db).Where("email = ?", email).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperr.New(apperr.ErrRecordNotFound, fmt.Sprintf("user with email `%s` not found", email))
		}
		slog.Error("error get account by email", "error", err)
		return nil, apperr.New(apperr.ErrInternalServer, "failed to get user")
	}

	domainUser := user.ToDomain()
	return &domainUser, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userModel := models.ToNewUserModel(user)
	err := gorm.G[models.UserModel](r.db).Create(ctx, &userModel)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, apperr.New(apperr.ErrEmailAlreadyExits, "invalid duplicate email")
		}
		slog.Error("error create user", "error", err)
		return nil, err
	}

	domainUser := userModel.ToDomain()
	return &domainUser, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain.User) error {
	modelUser := models.ToNewUserModel(user)
	rowAffected, err := gorm.G[models.UserModel](r.db).Updates(ctx, modelUser)
	if err != nil {
		return apperr.New(apperr.ErrInternalServer, "failed to update user")
	}

	if rowAffected == 0 {
		return apperr.New(apperr.ErrNoRecordAffected, "not user updated")
	}
	return nil
}
