package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type TimService struct {
	TimRepository repository.TimRepository
}

func NewTimService(timRepository repository.TimRepository) *TimService {
	return &TimService{TimRepository: timRepository}
}

func (service *TimService) GetAll() (*[]model.Tim, error) {
	teams, err := service.TimRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams were found"))
	}

	return &teams, nil
}

func (service *TimService) GetByID(id int) (*model.Tim, error) {
	tim, err := service.TimRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams with that id were found"))
	}

	return tim, nil
}
