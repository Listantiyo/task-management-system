package usecase

import (
	"context"
	"task-management-system/internal/delivery/dto"
	apperr "task-management-system/internal/delivery/error"
	"task-management-system/internal/domain"
	"task-management-system/internal/pkg/jwt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) Login(ctx context.Context, reqLogin dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.repo.GetUserByEmail(ctx, reqLogin.Email)
	if err != nil {
		return nil, err
	}
	if !user.ValidatePassword(reqLogin.Password) {
		return nil, apperr.New(apperr.ErrPasswordWrong, "invalid password")
	}

	tokenString, err := jwt.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	result := &dto.LoginResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Token:    tokenString,
	}

	return result, nil
}

func (u *userUsecase) Register(ctx context.Context, reqRegister dto.RegisterRequest) (*dto.RegisterResponse, error) {
	if _, err := u.repo.GetUserByEmail(ctx, reqRegister.Email); err != nil {
		return nil, err
	}

	newUser := domain.User{
		Email:    reqRegister.Email,
		Username: reqRegister.Username,
	}
	newUser.SetPassword(reqRegister.Password)
	userCreated, err := u.repo.CreateUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	result := &dto.RegisterResponse{
		ID:        userCreated.ID,
		Email:     userCreated.Email,
		Username:  userCreated.Username,
		CreatedAt: userCreated.CreatedAt,
	}
	return result, nil
}

// func (u *userUsecase) UpdateUser(ctx context.Context, reqUpdate dto.UpdateUserRequest) error
