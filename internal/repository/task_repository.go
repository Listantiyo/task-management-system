package repository

import (
	"context"
	"task-management-system/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(ctx context.Context, task *domain.Task) (*domain.Task, error) {
	return nil, nil
}

func (r *taskRepo) FindByID(ctx context.Context, id uuid.UUID) (*domain.Task, error)
func (r *taskRepo) FindAll(ctx context.Context, limit, offset int64) ([]domain.Task, int, error)
func (r *taskRepo) Update(ctx context.Context, id uuid.UUID, task *domain.Task) error
func (r *taskRepo) Delete(ctx context.Context, id uuid.UUID) error
