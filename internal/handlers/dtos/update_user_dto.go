package dtos

type UpdateUserDTO struct {
	Email    *string `json:"email" validate:"email"`
	Password *string `json:"password"`
}
