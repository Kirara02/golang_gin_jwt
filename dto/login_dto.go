package dto

type LoginDTO struct {
	Email    uint64 `json:"email" form:"email" binding:"required" validate:"email"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}
