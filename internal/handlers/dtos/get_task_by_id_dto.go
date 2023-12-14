package dtos

type GetTaskByIdDTO struct {
	Id string `json:"id" binding:"required"`
}
