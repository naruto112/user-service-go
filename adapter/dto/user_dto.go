package dto

import (
	"time"
	"user-service/adapter/request/entity"
	userEntity "user-service/core/domain/entity"
)

type UserDTO struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewUserDTO(u *entity.UserRequest) *userEntity.User {
	return &userEntity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}
