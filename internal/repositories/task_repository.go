package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
)

type TaskRepository struct{}

func NewTaskRepository() ports.TaskRepository {
	return &TaskRepository{}
}

func (t *TaskRepository) GetTasks() ([]*models.Task, error) {
	// TODO: implement
	return nil, nil
}

func (t *TaskRepository) CreateTask(*models.Task) error {
	// TODO: implement
	return nil
}

func (t *TaskRepository) GetTaskById(taskId uuid.UUID) (*models.Task, error) {
	// TODO: implement
	return nil, nil
}

func (t *TaskRepository) UpdateTask(taskId uuid.UUID, title *string, description *string, toDate *time.Time, completed *bool) error {
	// TODO: implement
	return nil
}

func (t *TaskRepository) DeleteTask(taskId uuid.UUID) error {
	// TODO: implement
	return nil
}
