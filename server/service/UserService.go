package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(UserRepository repository.UserRepository) *UserService {
	return &UserService{UserRepository: UserRepository}
}

func (service *UserService) GetAll() (*[]model.User, error) {
	users, err := service.UserRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users were found"))
	}

	return &users, nil
}

func (service *UserService) GetByID(id int) (*model.User, error) {
	user, err := service.UserRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users with that id were found"))
	}

	return user, nil
}

func (service *UserService) GetByEmail(email string) (*model.User, error) {
	user, err := service.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no users with that email were found"))
	}

	return user, nil
}