package dto

import (
	customType "task-management-system/internal/delivery/types"
)

type CreateTaskRequest struct {
	Title        string                `json:"title" binding:"required"`
	Status       customType.TaskStatus `json:"status" binding:"required,enum"`
	Descriptions string                `json:"descriptions"`
}

type UpdateTaskRequest struct {
	ID           string                `json:"task_id" binding:"required"`
	Title        string                `json:"title"`
	Status       customType.TaskStatus `json:"status" binding:"enum"`
	Descriptions string                `json:"descriptions"`
}

type TaskResponse struct {
	Tasks []any `json:"tasks"`
}
