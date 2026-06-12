package repository

import (
	"context"
	"errors"
	"task-management-system/internal/domain"
	"task-management-system/internal/models"

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

func (r *taskRepo) FindByID(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) (*domain.Task, error) {
	modelTask, err := gorm.G[models.TaskModel](r.db).Where(`user_id = ? AND id = ?`, userID, taskID).First(ctx)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	task := modelTask.ToDomain()
	return &task, nil
}
func (r *taskRepo) FindAll(ctx context.Context, limit, offset int64) ([]domain.Task, int, error)
func (r *taskRepo) Update(ctx context.Context, id uuid.UUID, task *domain.Task) error
func (r *taskRepo) Delete(ctx context.Context, id uuid.UUID) error
