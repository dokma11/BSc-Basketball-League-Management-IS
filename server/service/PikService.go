package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type PikService struct {
	PikRepository repository.PikRepository
}

func NewPikService(pikRepository repository.PikRepository) *PikService {
	return &PikService{PikRepository: pikRepository}
}

func (service *PikService) GetAll() (*[]model.Pik, error) {
	picks, err := service.PikRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks were found"))
	}

	return &picks, nil
}

func (service *PikService) GetByID(id int) (*model.Pik, error) {
	pick, err := service.PikRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks with that id were found"))
	}

	return pick, nil
}

func (service *PikService) GetAllByTeamID(teamId int) (*[]model.Pik, error) {
	picks, err := service.PikRepository.GetAllByTeamID(teamId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks with that team id were found"))
	}

	return &picks, nil
}

func (service *PikService) GetAllByYear(year string) (*[]model.Pik, error) {
	picks, err := service.PikRepository.GetAllByYear(year)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks in that year were found"))
	}

	return &picks, nil
}
