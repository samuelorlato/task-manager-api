package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/samuelorlato/task-manager-api/internal/configs"
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
	"github.com/samuelorlato/task-manager-api/pkg/validation"
)

type TaskService struct {
	repository ports.TaskRepository
}

func NewTaskService(repository ports.TaskRepository) ports.TaskUsecase {
	return &TaskService{
		repository: repository,
	}
}

func (t *TaskService) GetTasks() ([]*models.Task, error) {
	tasks, err := t.repository.GetTasks()
	if err != nil {
		// TODO: handle
	}

	return tasks, nil
}

func (t *TaskService) CreateTask(title string, description *string, toDate string) error {
	parsedToDate, err := time.Parse(configs.ToDateTaskLayout, toDate)
	if err != nil {
		// TODO: handle
	}

	task := models.NewTask(title, description, parsedToDate, false)

	err = validation.ValidateStruct(*task)
	if err != nil {
		// TODO: handle
	}

	err = t.repository.CreateTask(task)
	if err != nil {
		// TODO: handle
	}

	return nil
}

func (t *TaskService) GetTaskById(taskId string) (*models.Task, error) {
	taskIdUUID, err := uuid.FromBytes([]byte(taskId))
	if err != nil {
		// TODO: handle
	}

	task, err := t.repository.GetTaskById(taskIdUUID)
	if err != nil {
		// TODO: handle
	}

	return task, nil
}

func (t *TaskService) UpdateTask(taskId string, title *string, description *string, toDate *string, completed *bool) error {
	var parsedToDate time.Time

	if toDate != nil {
		parsed, err := time.Parse(configs.ToDateTaskLayout, *toDate)
		if err != nil {
			// TODO: handle
		}

		parsedToDate = parsed
	}

	taskIdUUID, err := uuid.FromBytes([]byte(taskId))
	if err != nil {
		// TODO: handle
	}

	err = t.repository.UpdateTask(taskIdUUID, title, description, &parsedToDate, completed)
	if err != nil {
		// TODO: handle
	}

	return nil
}

func (t *TaskService) DeleteTask(taskId string) error {
	taskIdUUID, err := uuid.FromBytes([]byte(taskId))
	if err != nil {
		// TODO: handle
	}

	err = t.repository.DeleteTask(taskIdUUID)
	if err != nil {
		// TODO: handle
	}

	return nil
}
