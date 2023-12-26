package dtos

type TaskDTO struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	ToDate      string    `json:"toDate"`
	Completed   bool      `json:"completed"`
	Tags        *[]string `json:"tags"`
}
