package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type IgracService struct {
	IgracRepository repository.IgracRepository
}

func NewIgracService(igracRepository repository.IgracRepository) *IgracService {
	return &IgracService{IgracRepository: igracRepository}
}

func (service *IgracService) GetAll() (*[]model.Igrac, error) {
	players, err := service.IgracRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no players were found"))
	}

	return &players, nil
}

func (service *IgracService) GetByID(id int) (*model.Igrac, error) {
	player, err := service.IgracRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no players with that id were found"))
	}

	return player, nil
}

func (service *IgracService) GetAllByTeamID(teamId int) (*[]model.Igrac, error) {
	players, err := service.IgracRepository.GetAllByTeamID(teamId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no players with that team id were found"))
	}

	return &players, nil
}
