package repository

import (
	"time"
	users "user-service/src/adapter/entity"

	"gorm.io/gorm"
)

type InterfaceUserRepository interface {
	CreateUser(user *users.User) error
	GetUser(id uint) (*users.User, error)
	UpdateUser(user *users.User) error
	DeleteUser(id uint) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *users.User) error {
	user.UpdateAt = time.Now()
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUser(id uint) (*users.User, error) {
	var user users.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetUserAll() ([]*users.User, error) {
	var users []*users.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) UpdateUser(user *users.User) (int64, error) {
	user.UpdateAt = time.Now()
	userCoy := *user
	userRow := r.db.Find(&userCoy)
	if userRow.RowsAffected != 0 {
		err := r.db.Model(user).Updates(user).Error
		if err != nil {
			return 0, err
		}
	}
	return userRow.RowsAffected, nil
}

func (r *UserRepository) DeleteUser(id uint) (int64, error) {
	var count int64
	r.db.Model(&users.User{}).Where("id = ?", id).Count(&count)
	err := r.db.Debug().Delete(&users.User{}, id).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
