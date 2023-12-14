package dtos

type DeleteTaskDTO struct {
	Id string `json:"id" binding:"required"`
}
