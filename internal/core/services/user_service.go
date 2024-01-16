package services

import (
	"github.com/samuelorlato/task-manager-api/internal/core/models"
	"github.com/samuelorlato/task-manager-api/internal/core/ports"
	"github.com/samuelorlato/task-manager-api/pkg/errors"
	"github.com/samuelorlato/task-manager-api/pkg/validation"
)

type UserService struct {
	repository        ports.UserRepository
	encryptionService ports.EncryptionService
}

func NewUserService(repository ports.UserRepository, encryptionService ports.EncryptionService) ports.UserUsecase {
	return &UserService{
		repository:        repository,
		encryptionService: encryptionService,
	}
}

func (u *UserService) GetUser(email string, password string) (*models.User, *errors.HTTPError) {
	user, err := u.repository.GetUser(email)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return nil, err
	}

	err = u.encryptionService.CompareHashAndPassword(user.Password, password)
	if err != nil {
		err := errors.NewValidationError(err)
		return nil, err
	}

	return user, nil
}

func (u *UserService) CreateUser(email string, password string) *errors.HTTPError {
	user := models.NewUser(email, password)

	err := validation.ValidateStruct(*user)
	if err != nil {
		err := errors.NewValidationError(err)
		return err
	}

	hashedPassword, err := u.encryptionService.HashPassword(user.Password)
	if err != nil {
		err := errors.NewGenericError(err)
		return err
	}

	user.Password = *hashedPassword

	err = u.repository.CreateUser(user)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return err
	}

	return nil
}

func (u *UserService) UpdateUser(loggedAsEmail string, email *string, password *string) *errors.HTTPError {
	if password != nil {
		hashedPassword, err := u.encryptionService.HashPassword(*password)
		if err != nil {
			err := errors.NewGenericError(err)
			return err
		}

		password = hashedPassword
	}

	err := u.repository.UpdateUser(loggedAsEmail, email, password)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return err
	}

	return nil
}

func (u *UserService) DeleteUser(email string) *errors.HTTPError {
	err := u.repository.DeleteUser(email)
	if err != nil {
		err := errors.NewRepositoryError(err)
		return err
	}

	return nil
}
