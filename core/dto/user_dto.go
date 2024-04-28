package dto

import (
	"time"
	userEntity "user-service/adapter/entity"
	"user-service/core/domain/entity"
)

type UserDTO struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewUserDTO(u *entity.User) *userEntity.User {
	return &userEntity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdateAt:  u.UpdateAt,
	}

}
