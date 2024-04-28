package services

import (
	"user-service/adapter/db"
	"user-service/adapter/repository"
	"user-service/core/domain/entity"
	"user-service/core/dto"
)

type UserServicesInterface interface {
	CreateUser(user *entity.User) error
	GetUser(id uint) (*entity.User, error)
	GetUserAll() ([]*entity.User, error)
	UpdateUser(user *entity.User) (int64, error)
	DeleteUser(id uint) (int64, error)
}

type UserServices struct {
	User *entity.User
}

func NewUserServices(user *entity.User) UserServicesInterface {
	return &UserServices{User: user}
}

func (s *UserServices) CreateUser(user *entity.User) error {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userDTO := dto.NewUserDTO(user)
	userRepository.CreateUser(userDTO)
	return nil
}

func (s *UserServices) GetUser(id uint) (*entity.User, error) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	user, err := userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdateAt,
	}, nil
}

func (s *UserServices) GetUserAll() ([]*entity.User, error) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	users, err := userRepository.GetUserAll()
	if err != nil {
		return nil, err
	}

	entityUsers := make([]*entity.User, len(users))
	for i, user := range users {
		entityUsers[i] = &entity.User{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdateAt:  user.UpdateAt,
		}
	}
	return entityUsers, nil
}

func (s *UserServices) UpdateUser(user *entity.User) (int64, error) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userDTO := dto.NewUserDTO(user)
	rows, err := userRepository.UpdateUser(userDTO)
	return rows, err
}

func (s *UserServices) DeleteUser(id uint) (int64, error) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	rowEffected, _ := userRepository.DeleteUser(id)
	return rowEffected, nil
}
