package models

import "time"

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ToDate      time.Time `json:"toDate"`
	Completed   bool      `json:"completed"`
}
