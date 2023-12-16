package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/samuelorlato/task-manager-api/internal/configs"
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
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

func (t *TaskService) GetTasks() ([]*models.Task, *errors.HTTPError) {
	tasks, err := t.repository.GetTasks()
	if err != nil {
		err := errors.NewRepositoryError(err)
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) CreateTask(title string, description *string, toDate string, tags *[]string) *errors.HTTPError {
	parsedToDate, err := time.Parse(configs.ToDateTaskLayout, toDate)
	if err != nil {
		err := errors.NewValidationError(err)
		return err
	}

	if tags != nil {
		if len(*tags) == 0 {
			tags = nil
		}
	}

	task := models.NewTask(title, description, parsedToDate, false, tags)

	err = validation.ValidateStruct(*task)
	if err != nil {
		err := errors.NewValidationError(err)
		return err
	}

	err = t.repository.CreateTask(task)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return err
	}

	return nil
}

func (t *TaskService) GetTaskById(taskId string) (*models.Task, *errors.HTTPError) {
	taskIdUUID, err := uuid.Parse(taskId)
	if err != nil {
		err := errors.NewValidationError(err)
		return nil, err
	}

	task, err := t.repository.GetTaskById(taskIdUUID)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return nil, err
	}

	return task, nil
}

func (t *TaskService) UpdateTask(taskId string, title *string, description *string, toDate *string, completed *bool, tags *[]string) *errors.HTTPError {
	var parsedToDate time.Time

	if toDate != nil && *toDate != "" {
		parsed, err := time.Parse(configs.ToDateTaskLayout, *toDate)
		if err != nil {
			err := errors.NewValidationError(err)
			return err
		}

		parsedToDate = parsed
	}

	taskIdUUID, err := uuid.Parse(taskId)
	if err != nil {
		err := errors.NewValidationError(err)
		return err
	}

	if tags != nil {
		if len(*tags) == 0 {
			tags = nil
		}
	}

	err = t.repository.UpdateTask(taskIdUUID, title, description, &parsedToDate, completed, tags)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return err
	}

	return nil
}

func (t *TaskService) DeleteTask(taskId string) *errors.HTTPError {
	taskIdUUID, err := uuid.Parse(taskId)
	if err != nil {
		err := errors.NewValidationError(err)
		return err
	}

	err = t.repository.DeleteTask(taskIdUUID)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return err
	}

	return nil
}
