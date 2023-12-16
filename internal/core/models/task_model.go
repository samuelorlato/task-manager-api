package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `validate:"required,uuid"`
	Title       string    `validate:"required"`
	Description *string
	ToDate      time.Time `validate:"required"`
	Completed   *bool     `validate:"required"`
	Tags        *[]string
}

func NewTask(title string, description *string, toDate time.Time, completed bool, tags *[]string) *Task {
	return &Task{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		ToDate:      toDate,
		Completed:   &completed,
		Tags:        tags,
	}
}
