package dtos

type UpdateUserDTO struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}
