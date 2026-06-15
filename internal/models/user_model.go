package models

import (
	"task-management-system/internal/domain"
	"time"

	"github.com/google/uuid"
)

// Model
type UserModel struct {
	ID           uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create"`
	Email        string      `gorm:"type:varchar(255);not null;unique"`
	Username     string      `gorm:"type:varchar(100);not null;unique"`
	PasswordHash string      `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time   `gorm:"type:timestamptz;default:CURRENT_TIMESAMP;<-:create"`
	UpdatedAt    time.Time   `gorm:"type:timestamptz;autoUpdateTime"`
	Tasks        []TaskModel `gorm:"foreignKey:UserID"`
}

func (u *UserModel) ToDomain() domain.User {
	return domain.User{
		ID:           u.ID,
		Email:        u.Email,
		Username:     u.Username,
		PasswordHash: u.PasswordHash,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func ToNewUserModel(user *domain.User) UserModel {
	return UserModel{
		Email:        user.Email,
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
	}
}

func (UserModel) TableName() string {
	return "users"
}
