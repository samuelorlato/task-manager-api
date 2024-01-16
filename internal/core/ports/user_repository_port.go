package ports

import (
	"github.com/samuelorlato/task-manager-api/internal/core/models"
)

type UserRepository interface {
	GetUser(email string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(loggedAsEmail string, email *string, password *string) error
	DeleteUser(email string) error
}
