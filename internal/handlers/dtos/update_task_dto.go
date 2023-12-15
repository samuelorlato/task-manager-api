package dtos

type UpdateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ToDate      string `json:"toDate"`
	Completed   bool   `json:"completed"`
}
