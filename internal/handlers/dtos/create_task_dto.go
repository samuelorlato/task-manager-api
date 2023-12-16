package dtos

type CreateTaskDTO struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	ToDate      string   `json:"toDate" binding:"required"`
	Tags        []string `json:"tags"`
}
