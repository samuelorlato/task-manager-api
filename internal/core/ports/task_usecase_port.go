package ports

import (
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
)

type TaskUsecase interface {
	GetTasks(email string) ([]*models.Task, *errors.HTTPError)
	CreateTask(email string, title string, description *string, toDate string, tasks *[]string) (*string, *errors.HTTPError)
	GetTaskById(email string, taskId string) (*models.Task, *errors.HTTPError)
	UpdateTask(email string, taskId string, title *string, description *string, toDate *string, completed *bool, tasks *[]string) *errors.HTTPError
	DeleteTask(email string, taskId string) *errors.HTTPError
}
