package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

var _ callbacks.BeforeCreateInterface = (*UserModel)(nil)

// Model
type UserModel struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Email        string     `gorm:"type:varchar(50);not null;unique"`
	Username     string     `gorm:"type:varchar(100);not null"`
	PasswordHash string     `gorm:"type:varchar(100);not null"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime"`
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	now := time.Now().Truncate(time.Second)
	u.ID = newUUID
	u.CreatedAt = now
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now().Truncate(time.Second)
	u.UpdatedAt = &now
	tx.Statement.SetColumn("updated_at", now)
	return nil
}
