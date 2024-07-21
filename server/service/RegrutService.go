package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type RegrutService struct {
	RegrutRepository repository.RegrutRepository
}

func NewRegrutService(regrutRepository repository.RegrutRepository) *RegrutService {
	return &RegrutService{RegrutRepository: regrutRepository}
}

func (service *RegrutService) GetAll() (*[]model.Regrut, error) {
	recruits, err := service.RegrutRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no recruits were found"))
	}

	return &recruits, nil
}

func (service *RegrutService) GetByID(id int) (*model.Regrut, error) {
	recruit, err := service.RegrutRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no recruits with that id were found"))
	}

	return recruit, nil
}

func (service *RegrutService) Create(recruit *model.Regrut) error {
	err := service.RegrutRepository.Create(recruit)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no recruits were created"))
		return err
	}
	return nil
}

func (service *RegrutService) Update(recruit *model.Regrut) error {
	err := service.RegrutRepository.Update(recruit)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no recruits were updated"))
		return err
	}
	return nil
}
