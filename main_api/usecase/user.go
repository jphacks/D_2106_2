package usecase

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func (uc *UserUsecase) RegisterNewUser(userName string, password string) (int, error) {
	user := &domain.User{
		Name:     userName,
		Password: password,
	}

	userId, err := uc.UserRepo.CreateUser(user)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (uc *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	users, err := uc.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUsecase) GetUserById(userId int) (*domain.User, error) {
	user, err := uc.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUsecase) CheckUserExist(username string) (bool, error) {
	_, err := uc.UserRepo.GetUserByName(username)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (uc *UserUsecase) Login(name string, password string) (*domain.User, error) {
	user, err := uc.UserRepo.GetUserByNameAndPasssword(name, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
