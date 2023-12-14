package ports

import "github.com/samuelorlato/task-manager-api/internal/core/models"

type TaskUsecase interface {
	GetTasks() ([]*models.Task, error)
	CreateTask(title string, description *string, toDate string) error
	GetTaskById(taskId string) (*models.Task, error)
	UpdateTask(taskId string, title *string, description *string, toDate *string, completed *bool) error
	DeleteTask(taskId string) error
}
