package dtos

type UpdateTaskDTO struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ToDate      string `json:"toDate"`
	Completed   bool   `json:"completed"`
}
