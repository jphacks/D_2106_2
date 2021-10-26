package repository

import "github.com/jphacks/D_2106_2/domain"

type UserRepository interface {
	CreateUser(user *domain.User) (int, error)
	GetAllUsers() ([]*domain.User, error)
	GetUserById(userId int) (*domain.User, error)
	GetUserByDeviceId(device_id string) (*domain.User, error)
}
