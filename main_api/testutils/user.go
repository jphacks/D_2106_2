package testutils

import "github.com/jphacks/D_2106_2/domain"

type FakeUserRepository struct {
	FakeCreateUser  func(user *domain.User) (string, error)
	FakeGetAllUsers func() ([]*domain.User, error)
	FakeGetUserById func(userId string) (*domain.User, error)
}

func (repo FakeUserRepository) CreateUser(user *domain.User) (string, error) {
	return repo.FakeCreateUser(user)
}

func (repo FakeUserRepository) GetAllUsers() ([]*domain.User, error) {
	return repo.FakeGetAllUsers()
}

func (repo FakeUserRepository) GetUserById(userId string) (*domain.User, error) {
	return repo.FakeGetUserById(userId)
}
