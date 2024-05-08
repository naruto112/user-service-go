package dto

import (
	"time"
	"user-service/src/adapter/request"
	"user-service/src/adapter/response"
	userEntity "user-service/src/core/domain/entity"
)

type UserDTO struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewUserDTORequest(u *request.UserRequest) *userEntity.User {
	return &userEntity.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}

func NewUserDTOLoginRequest(u *request.UserRequest) *userEntity.User {
	return &userEntity.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

func NewUserDTOResponse(u *userEntity.User) *response.UserResponse {
	return &response.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdateAt:  u.UpdateAt,
	}
}

func NewUserDTOResponseList(users []*userEntity.User) []response.UserResponse {
	userResponses := make([]response.UserResponse, len(users))
	for i, u := range users {
		userResponses[i] = response.UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdateAt:  u.UpdateAt,
		}
	}
	return userResponses
}
