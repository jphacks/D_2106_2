package testutils

import "github.com/jphacks/D_2106_2/domain"

type FakeUserRepository struct {
	FakeCreateUser        func(user *domain.User) (int, error)
	FakeGetAllUsers       func() ([]*domain.User, error)
	FakeGetUserById       func(userId int) (*domain.User, error)
	FakeGetUserByDeviceId func(deviceId string) (*domain.User, error)
}

func (repo FakeUserRepository) CreateUser(user *domain.User) (int, error) {
	return repo.FakeCreateUser(user)
}

func (repo FakeUserRepository) GetAllUsers() ([]*domain.User, error) {
	return repo.FakeGetAllUsers()
}

func (repo FakeUserRepository) GetUserById(userId int) (*domain.User, error) {
	return repo.FakeGetUserById(userId)
}

func (repo FakeUserRepository) GetUserByDeviceId(deviceId string) (*domain.User, error) {
	return repo.FakeGetUserByDeviceId(deviceId)
}
