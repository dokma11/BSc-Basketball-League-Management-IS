package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type TeamService struct {
	TeamRepository repository.TeamRepository
}

func NewTeamService(TeamRepository repository.TeamRepository) *TeamService {
	return &TeamService{TeamRepository: TeamRepository}
}

func (service *TeamService) GetAll() (*[]model.Team, error) {
	teams, err := service.TeamRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams were found"))
	}

	return &teams, nil
}

func (service *TeamService) GetByID(id int) (*model.Team, error) {
	team, err := service.TeamRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams with that id were found"))
	}

	return team, nil
}
