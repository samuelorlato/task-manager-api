package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `validate:"required,uuid"`
	Email    string    `validate:"required,email"`
	Password string    `validate:"required"`
}

func NewUser(email string, password string) *User {
	return &User{
		Id:       uuid.New(),
		Email:    email,
		Password: password,
	}
}
