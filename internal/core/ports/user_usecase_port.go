package ports

import (
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
)

type UserUsecase interface {
	GetUser(email string, password string) (*models.User, *errors.HTTPError)
	CreateUser(email string, password string) *errors.HTTPError
	UpdateUser(loggedAsEmail string, email *string, password *string) *errors.HTTPError
	DeleteUser(email string) *errors.HTTPError
}
