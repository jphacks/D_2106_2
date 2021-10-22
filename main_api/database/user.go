package database

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	return &UserRepository{sqlHandler}
}

func (repo *UserRepository) CreateUser(user *domain.User) (int, error) {
	result := repo.SqlHandler.Conn.Create(&user)
	if err := result.Error; err != nil {
		return -1, err
	}

	return user.Id, nil
}

func (repo *UserRepository) GetAllUsers() ([]*domain.User, error) {
	users := []*domain.User{}
	result := repo.SqlHandler.Conn.Find(&users)
	if err := result.Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (repo *UserRepository) GetUserById(userId int) (*domain.User, error) {
	user := &domain.User{}
	result := repo.SqlHandler.Conn.Where("ID = ?", userId).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserByName(name string) (*domain.User, error) {
	user := &domain.User{}
	result := repo.SqlHandler.Conn.Where("name = ?", name).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepository) GetUserByNameAndPasssword(name string, password string) (*domain.User, error) {
	user := &domain.User{}
	result := repo.SqlHandler.Conn.Where("name = ? and password = ?", name, password).First(&user)
	if err := result.Error; err != nil {
		return nil, err
	}

	return user, nil
}
