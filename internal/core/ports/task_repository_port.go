package ports

import (
	"time"

	"github.com/google/uuid"
	"github.com/samuelorlato/task-manager-api/internal/core/models"
)

type TaskRepository interface {
	GetTasks(email string) ([]*models.Task, error)
	CreateTask(task *models.Task) (*uuid.UUID, error)
	GetTaskById(email string, taskId uuid.UUID) (*models.Task, error)
	UpdateTask(email string, taskId uuid.UUID, title *string, description *string, toDate *time.Time, completed *bool, tags *[]string) error
	DeleteTask(email string, taskId uuid.UUID) error
}
