package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type PlayerService struct {
	PlayerRepository repository.PlayerRepository
}

func NewPlayerService(PlayerRepository repository.PlayerRepository) *PlayerService {
	return &PlayerService{PlayerRepository: PlayerRepository}
}

func (service *PlayerService) GetAll() (*[]model.Player, error) {
	players, err := service.PlayerRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no players were found"))
	}

	return &players, nil
}

func (service *PlayerService) GetByID(id int) (*model.Player, error) {
	player, err := service.PlayerRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no players with that id were found"))
	}

	return player, nil
}

func (service *PlayerService) GetAllByTeamID(teamId int) (*[]model.Player, error) {
	players, err := service.PlayerRepository.GetAllByTeamID(teamId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no players with that team id were found"))
	}

	return &players, nil
}

func (service *PlayerService) GetAllAvailableByTeamID(teamId int) (*[]model.Player, error) {
	players, err := service.PlayerRepository.GetAllAvailableByTeamID(teamId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no available players with that team id were found"))
	}

	return &players, nil
}

func (service *PlayerService) Update(player *model.Player) error {
	err := service.PlayerRepository.Update(player)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no players were updated"))
		return err
	}
	return nil
}
