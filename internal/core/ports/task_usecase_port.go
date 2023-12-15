package ports

import (
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
)

type TaskUsecase interface {
	GetTasks() ([]*models.Task, *errors.HTTPError)
	CreateTask(title string, description *string, toDate string) *errors.HTTPError
	GetTaskById(taskId string) (*models.Task, *errors.HTTPError)
	UpdateTask(taskId string, title *string, description *string, toDate *string, completed *bool) *errors.HTTPError
	DeleteTask(taskId string) *errors.HTTPError
}
