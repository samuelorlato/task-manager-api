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

type FirestoreRepository struct {
	firestoreClient *firestore.Client
}

func NewFirestoreRepository(firestoreClient *firestore.Client) ports.UserAndTaskRepository {
	return &FirestoreRepository{
		firestoreClient: firestoreClient,
	}
}

func (f *FirestoreRepository) GetTasks() ([]*models.Task, error) {
	docs, err := f.firestoreClient.Collection(configs.FirestoreTasksCollectionName).Documents(context.Background()).GetAll()
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

func (f *FirestoreRepository) CreateTask(task *models.Task) error {
	_, err := f.firestoreClient.Collection(configs.FirestoreTasksCollectionName).Doc(task.Id.String()).Set(context.Background(), map[string]interface{}{
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

func (f *FirestoreRepository) GetTaskById(taskId uuid.UUID) (*models.Task, error) {
	// TODO: implement
	return nil, nil
}

func (f *FirestoreRepository) UpdateTask(taskId uuid.UUID, title *string, description *string, toDate *time.Time, completed *bool, tags *[]string) error {
	// TODO: implement
	return nil
}

func (f *FirestoreRepository) DeleteTask(taskId uuid.UUID) error {
	// TODO: implement
	return nil
}

func (f *FirestoreRepository) GetUser(email string) (*models.User, error) {
	doc, err := f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Where("email", "==", email).Documents(context.Background()).Next()
	if err != nil {
		return nil, err
	}

	var user models.User
	err = doc.DataTo(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (f *FirestoreRepository) CreateUser(user *models.User) error {
	_, err := f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Doc(user.Id.String()).Set(context.Background(), map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	})

	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreRepository) UpdateUser(email *string, password *string) error {
	// TODO: implement
	return nil
}

func (f *FirestoreRepository) DeleteUser(email string) error {
	// TODO: implement
	return nil
}
