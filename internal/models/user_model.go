package models

import (
	"time"

	"github.com/google/uuid"
)

// Model
type UserModel struct {
	ID           uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string      `gorm:"type:varchar(255);not null;unique"`
	Username     string      `gorm:"type:varchar(100);not null;unique"`
	PasswordHash string      `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time   `gorm:"autoCreateTime"`
	UpdatedAt    time.Time   `gorm:"autoUpdateTime"`
	Tasks        []TaskModel `gorm:"foreignKey:UserID"`
}

func (UserModel) TableName() string {
	return "users"
}
