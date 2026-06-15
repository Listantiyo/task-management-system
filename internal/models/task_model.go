package models

import (
	customType "task-management-system/internal/delivery/types"
	"task-management-system/internal/domain"
	"time"

	"github.com/google/uuid"
)

type TaskModel struct {
	ID          uuid.UUID             `gorm:"type:uuid;default:gen_random_uuid();primaryKey;<-:create"`
	UserID      uuid.UUID             `gorm:"type:uuid;not null;index;<-:create"`
	Title       string                `gorm:"type:varchar(100);not null"`
	Description *string               `gorm:"type:text"`
	Status      customType.TaskStatus `gorm:"type:task_status;default:'PENDING';not null"`
	CreatedAt   time.Time             `gorm:"type:timestamptz;default:CURRENT_TIMESTAMP;<-:create"`
	UpdatedAt   time.Time             `gorm:"type:timestamptz;autoUpdateTime"`
	User        UserModel             `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (TaskModel) TableName() string {
	return "tasks"
}

func ToTaskModel(task *domain.Task) TaskModel {
	return TaskModel{
		UserID:      task.UserID,
		Title:       task.Title,
		Description: &task.Description,
		Status:      task.Status,
	}
}

func (m *TaskModel) ToDomain() domain.Task {
	return domain.Task{
		ID:          m.ID,
		UserID:      m.UserID,
		Title:       m.Title,
		Description: *m.Description,
		Status:      m.Status,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
