package repository

import "github.com/jphacks/D_2106_2/domain"

type UserRepository interface {
	CreateUser(user *domain.User) (string, error)
	GetAllUsers() ([]*domain.User, error)
	GetUserById(userId string) (*domain.User, error)
}
