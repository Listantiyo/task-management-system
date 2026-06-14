package repository

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	apperr "task-management-system/internal/delivery/error"
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
	modelTask := models.ToTaskModel(task)
	err := gorm.G[models.TaskModel](r.db).Create(ctx, &modelTask)
	if err != nil {
		slog.Error("failed create new task", "error", err)
		return nil, apperr.New(apperr.ErrInternalServer, "failed create new task")
	}
	domainTask := modelTask.ToDomain()
	return &domainTask, nil
}

func (r *taskRepo) FindByID(ctx context.Context, userID, taskID uuid.UUID) (*domain.Task, error) {
	modelTask, err := gorm.G[models.TaskModel](r.db).Where(`user_id = ? AND id = ?`, userID, taskID).First(ctx)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, apperr.New(apperr.ErrRecordNotFound, fmt.Sprintf("cannot found task with id: %s", taskID.String()))
		}
		slog.Error("failed to find task", "error", err)
		return nil, apperr.New(apperr.ErrInternalServer, "failed to find task")
	}

	taskDomain := modelTask.ToDomain()
	return &taskDomain, nil
}

func (r *taskRepo) FindAll(ctx context.Context, userID uuid.UUID, limit, offset int64) ([]domain.Task, int64, error) {
	modelTasks, err := gorm.G[models.TaskModel](r.db).Where("user_id = ?", userID).Limit(int(limit)).Offset(int(offset)).Find(ctx)
	if err != nil {
		slog.Error("failed find tasks", "error", err)
		return nil, 0, apperr.New(apperr.ErrInternalServer, "failed find tasks")
	}

	count := int64(len(modelTasks))
	var domainTasks = make([]domain.Task, count)
	if count > 0 {
		for i, task := range modelTasks {
			domainTasks[i] = task.ToDomain()
		}
		return domainTasks, count, nil
	}

	return domainTasks, count, nil
}

func (r *taskRepo) Update(ctx context.Context, userID, taskID uuid.UUID, task *domain.Task) error {
	modelTask := models.ToTaskModel(task)
	rowAffected, err := gorm.G[models.TaskModel](r.db).Where("user_id = ? AND id = ?", userID, taskID).Updates(ctx, modelTask)
	if err != nil {
		slog.Error("failed update task", "error", err)
		return apperr.New(apperr.ErrInternalServer, "failed to update task")
	}

	if rowAffected == 0 {
		return apperr.New(apperr.ErrNoRecordAffected, fmt.Sprintf("no data updated with id: %d", taskID))
	}

	return nil
}

func (r *taskRepo) Delete(ctx context.Context, userID uuid.UUID, taskID uuid.UUID) error {
	rowAffected, err := gorm.G[models.TaskModel](r.db).Where("id = ? AND user_id = ?", taskID, userID).Delete(ctx)
	if err != nil {
		return apperr.New(apperr.ErrInternalServer, "failed to delete task")
	}

	if rowAffected == 0 {
		return apperr.New(apperr.ErrNoRecordAffected, fmt.Sprintf("no data deleted with id: %d", taskID))
	}

	return nil
}
