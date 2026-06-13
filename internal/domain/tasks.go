package domain

import (
	"context"
	customType "task-management-system/internal/delivery/types"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Title        string
	Descriptions string
	Status       customType.TaskStatus
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Repo Interface
type TaskRepository interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	FindByID(ctx context.Context, userID uuid.UUID, taksID uuid.UUID) (*Task, error)
	FindAll(ctx context.Context, limit, offset int64) ([]Task, int, error)
	Update(ctx context.Context, id uuid.UUID, task *Task) error
	Delete(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) error
}

// Usecase Interface
type TaskUsecase interface {
}
