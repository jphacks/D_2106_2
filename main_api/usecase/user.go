package usecase

import (
	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func (uc *UserUsecase) RegisterNewUser(
	userName string,
	deviceId string,
	introduction string,
) (int, error) {
	user := &domain.User{
		Name:         userName,
		DeviceId:     deviceId,
		Introduction: introduction,
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

func (uc *UserUsecase) GetUserByDeviceId(device_id string) (*domain.User, error) {
	user, err := uc.UserRepo.GetUserByDeviceId(device_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
