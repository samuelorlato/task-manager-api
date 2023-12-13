package services

import (
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
)

type TaskService struct{}

func NewTaskService() ports.TaskUsecase {
	return &TaskService{}
}

func (*TaskService) CreateTask(title string, description *string, toDate string) error {
	// TODO: implement
}

func (*TaskService) DeleteTask(taskId string) error {
	// TODO: implement
}

func (*TaskService) GetTaskById(taskId string) (*models.Task, error) {
	// TODO: implement
}

func (*TaskService) GetTasks() []*models.Task {
	// TODO: implement
}

func (*TaskService) UpdateTask(taskId string, title *string, description *string, toDate *string, completed *bool) error {
	// TODO: implement
}
