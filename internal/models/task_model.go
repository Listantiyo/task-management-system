package models

import (
	customType "task-management-system/internal/delivery/types"
	"time"

	"github.com/google/uuid"
)

type TaskModel struct {
	ID          uuid.UUID             `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID             `gorm:"type:uuid;not null;index"`
	Title       string                `gorm:"type:varchar(100);not null"`
	Description string                `gorm:"type:text"`
	Status      customType.TaskStatus `gorm:"type:task_status;default:'PENDING'; not null"`
	CreatedAt   time.Time             `gorm:"autoCreateTime"`
	UpdatedAt   time.Time             `gorm:"autoUpdateTime"`
	User        UserModel             `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (TaskModel) TableName() string {
	return "tasks"
}
