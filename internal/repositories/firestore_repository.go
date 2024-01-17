package repositories

import (
	"context"
	"encoding/json"
	"errors"
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

func (f *FirestoreRepository) GetTasks(email string) ([]*models.Task, error) {
	docs, err := f.firestoreClient.Collection(configs.FirestoreTasksCollectionName).Where("from", "==", email).Documents(context.Background()).GetAll()
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

func (f *FirestoreRepository) CreateTask(task *models.Task) (*uuid.UUID, error) {
	_, err := f.firestoreClient.Collection(configs.FirestoreTasksCollectionName).Doc(task.Id.String()).Set(context.Background(), map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"toDate":      task.ToDate,
		"completed":   task.Completed,
		"tags":        task.Tags,
		"from":        task.From,
	})

	if err != nil {
		return nil, err
	}

	return &task.Id, nil
}

func (f *FirestoreRepository) GetTaskById(email string, taskId uuid.UUID) (*models.Task, error) {
	tasks, err := f.GetTasks(email)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.Id == taskId {
			return task, nil
		}
	}

	return nil, errors.New("Task not found")
}

func (f *FirestoreRepository) UpdateTask(email string, taskId uuid.UUID, title *string, description *string, toDate *time.Time, completed *bool, tags *[]string) error {
	tasks, err := f.GetTasks(email)
	if err != nil {
		return err
	}

	var foundTask *models.Task
	for _, task := range tasks {
		if task.Id == taskId {
			foundTask = task
			break
		}
	}

	updates := map[string]interface{}{}
	if title != nil {
		updates["title"] = title
	}
	if description != nil {
		updates["description"] = description
	}
	if toDate != nil {
		updates["toDate"] = toDate
	}
	if completed != nil {
		updates["completed"] = completed
	}
	if tags != nil {
		updates["tags"] = tags
	}

	firestoreUpdates := []firestore.Update{}
	for key, value := range updates {
		firestoreUpdates = append(firestoreUpdates, firestore.Update{
			Path:  key,
			Value: value,
		})
	}

	_, err = f.firestoreClient.Collection(configs.FirestoreTasksCollectionName).Doc(foundTask.Id.String()).Update(context.Background(), firestoreUpdates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreRepository) DeleteTask(email string, taskId uuid.UUID) error {
	tasks, err := f.GetTasks(email)
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if task.Id == taskId {
			_, err := f.firestoreClient.Collection(configs.FirestoreTasksCollectionName).Doc(taskId.String()).Delete(context.Background())
			if err != nil {
				return err
			}

			return nil
		}
	}

	return errors.New("Task not found")
}

func (f *FirestoreRepository) GetUser(email string) (*models.User, error) {
	doc, err := f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Where("email", "==", email).Documents(context.Background()).Next()
	if err != nil {
		err := errors.New("User not found")
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
	_, err := f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Where("email", "==", user.Email).Documents(context.Background()).Next()
	if err == nil {
		return errors.New("Email used on another account")
	}

	_, err = f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Doc(user.Id.String()).Set(context.Background(), map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
	})

	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreRepository) UpdateUser(loggedAsEmail string, email *string, password *string) error {
	users, err := f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Where("email", "==", loggedAsEmail).Documents(context.Background()).GetAll()
	if err != nil {
		return err
	}

	userId := users[0].Ref.ID

	updates := map[string]interface{}{}
	if email != nil {
		updates["email"] = email
	}
	if password != nil {
		updates["password"] = password
	}

	firestoreUpdates := []firestore.Update{}
	for key, value := range updates {
		firestoreUpdates = append(firestoreUpdates, firestore.Update{
			Path:  key,
			Value: value,
		})
	}

	if email != nil {
		tasksToChange, err := f.GetTasks(loggedAsEmail)
		if err != nil {
			return err
		}

		for _, task := range tasksToChange {
			taskUpdate := []firestore.Update{
				{
					Path: "from",
					Value: email,
				},
			}
			_, err = f.firestoreClient.Collection(configs.FirestoreTasksCollectionName).Doc(task.Id.String()).Update(context.Background(), taskUpdate)
			if err != nil {
				return err
			}
		}
	}

	_, err = f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Doc(userId).Update(context.Background(), firestoreUpdates)
	if err != nil {
		return err
	}

	return nil
}

func (f *FirestoreRepository) DeleteUser(email string) error {
	users, err := f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Where("email", "==", email).Documents(context.Background()).GetAll()
	if err != nil {
		return err
	}

	userId := users[0].Ref.ID

	_, err = f.firestoreClient.Collection(configs.FirestoreUsersCollectionName).Doc(userId).Delete(context.Background())
	if err != nil {
		return err
	}

	return nil
}
