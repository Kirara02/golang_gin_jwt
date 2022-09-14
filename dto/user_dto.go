package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    uint64 `json:"email" form:"email" binding:"required"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}

type UserCreateDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    uint64 `json:"email" form:"email" binding:"required" validate:"email"`
	Password uint64 `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
}
