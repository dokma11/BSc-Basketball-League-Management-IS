package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type ZaposleniService struct {
	ZaposleniRepository repository.ZaposleniRepository
}

func NewZaposleniService(zaposleniRepository repository.ZaposleniRepository) *ZaposleniService {
	return &ZaposleniService{ZaposleniRepository: zaposleniRepository}
}

func (service *ZaposleniService) GetAll() (*[]model.Zaposleni, error) {
	employees, err := service.ZaposleniRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no employees were found"))
	}

	return &employees, nil
}

func (service *ZaposleniService) GetByID(id int) (*model.Zaposleni, error) {
	employee, err := service.ZaposleniRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no employees with that id were found"))
	}

	return employee, nil
}
