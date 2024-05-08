package services

import (
	"errors"
	"user-service/src/adapter/code"
	"user-service/src/adapter/db"
	"user-service/src/adapter/repository"
	"user-service/src/core/domain/entity"
	"user-service/src/core/dto"
)

type UserServicesInterface interface {
	CreateUser(user *entity.User) error
	LoginUser(user *entity.User) (map[string]interface{}, error)
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

func (s *UserServices) LoginUser(user *entity.User) (map[string]interface{}, error) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userDTO := dto.NewUserDTO(user)
	u, err := userRepository.GetUser(userDTO)

	if err != nil {
		return nil, errors.New("user not found")
	}

	isValid := code.DecryptHashPassword(&user.Password, &userDTO.Password)

	if !isValid {
		return nil, errors.New("invalid credentials")
	}

	return code.GenerateToken(&u.Name), err

}

func (s *UserServices) CreateUser(user *entity.User) error {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	userDTO := dto.NewUserDTO(user)
	password := code.HashPassword(&userDTO.Password)
	userDTO.Password = password
	userRepository.CreateUser(userDTO)
	return nil
}

func (s *UserServices) GetUser(id uint) (*entity.User, error) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	user, err := userRepository.GetUserByID(id)
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
	password := code.HashPassword(&userDTO.Password)
	userDTO.Password = password
	rows, err := userRepository.UpdateUser(userDTO)
	return rows, err
}

func (s *UserServices) DeleteUser(id uint) (int64, error) {
	userRepository := repository.NewUserRepository(db.Mysqlconnection())
	rowEffected, _ := userRepository.DeleteUser(id)
	return rowEffected, nil
}
