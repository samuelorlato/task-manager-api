package ports

import (
	"time"

	"github.com/google/uuid"
	"github.com/samuelorlato/task-manager-api/internal/core/models"
)

type TaskRepository interface {
	GetTasks() ([]*models.Task, error)
	CreateTask(*models.Task) error
	GetTaskById(taskId uuid.UUID) (*models.Task, error)
	UpdateTask(taskId uuid.UUID, title *string, description *string, toDate *time.Time, completed *bool) error
	DeleteTask(taskId uuid.UUID) error
}
