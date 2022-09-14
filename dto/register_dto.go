package dto

type Register struct {
	Name     string `json:"name" form:"name" binding:"required" validate:"min:1"`
	Email    uint64 `json:"email" form:"email" binding:"required" validate:"email"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}
