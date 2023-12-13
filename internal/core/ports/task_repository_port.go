package ports

import "github.com/samuelorlato/task-manager-api/internal/core/models"

type TaskRepository interface {
	CreateTask(title string, description *string, toDate string) error
	GetTask(taskId string) (*models.Task, error)
	UpdateTask(taskId string, title *string, description *string, toDate *string, completed *bool) error
	DeleteTask(taskId string) error
}
