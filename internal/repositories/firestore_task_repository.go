package repositories

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"github.com/samuelorlato/task-manager-api/internal/configs"
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
)

type FirestoreTaskRepository struct {
	firestoreClient *firestore.Client
}

func NewFirestoreTaskRepository(firestoreClient *firestore.Client) ports.TaskRepository {
	return &FirestoreTaskRepository{
		firestoreClient: firestoreClient,
	}
}

func (t *FirestoreTaskRepository) GetTasks() ([]*models.Task, error) {
	docs, err := t.firestoreClient.Collection(configs.FirestoreCollectionName).Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}

	var tasks []*models.Task
	for _, doc := range docs {
		var taskData map[string]interface{}
		if err := doc.DataTo(&taskData); err != nil {
			return nil, err
		}

		jsonBody, err := json.Marshal(taskData)
		if err != nil {
			return nil, err
		}

		var task models.Task
		if err := json.Unmarshal(jsonBody, &task); err != nil {
			return nil, err
		}
		task.Id = uuid.MustParse(doc.Ref.ID)

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (t *FirestoreTaskRepository) CreateTask(task *models.Task) error {
	_, err := t.firestoreClient.Collection(configs.FirestoreCollectionName).Doc(task.Id.String()).Set(context.Background(), map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"toDate":      task.ToDate,
		"completed":   task.Completed,
		"tags":        task.Tags,
	})

	if err != nil {
		return err
	}

	return nil
}

func (t *FirestoreTaskRepository) GetTaskById(taskId uuid.UUID) (*models.Task, error) {
	// TODO: implement
	return nil, nil
}

func (t *FirestoreTaskRepository) UpdateTask(taskId uuid.UUID, title *string, description *string, toDate *time.Time, completed *bool, tags *[]string) error {
	// TODO: implement
	return nil
}

func (t *FirestoreTaskRepository) DeleteTask(taskId uuid.UUID) error {
	// TODO: implement
	return nil
}
