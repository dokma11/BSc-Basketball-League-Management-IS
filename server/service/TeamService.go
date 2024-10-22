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

func (service *TeamService) GetByUserID(userID int) (*model.Team, error) {
	team, err := service.TeamRepository.GetByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams with that user id were found"))
	}
	return team, nil
}

func (service *TeamService) GetPlayerTradeDestination(tradeSubjectID int) (*model.Team, error) {
	team, err := service.TeamRepository.GetPlayerTradeDestination(tradeSubjectID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams with that trade subject id were found"))
	}
	return team, nil
}

func (service *TeamService) GetPickTradeDestination(tradeSubjectID int) (*model.Team, error) {
	team, err := service.TeamRepository.GetPickTradeDestination(tradeSubjectID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams with that trade subject id were found"))
	}
	return team, nil
}

func (service *TeamService) GetDraftRightsTradeDestination(tradeSubjectID int) (*model.Team, error) {
	team, err := service.TeamRepository.GetDraftRightsTradeDestination(tradeSubjectID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no teams with that trade subject id were found"))
	}
	return team, nil
}

func (service *TeamService) GetWishlistByTeamID(teamID int) ([]model.WishlistAsset, error) {
	wishlist, err := service.TeamRepository.GetWishlistByTeamID(teamID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no wishlists with that team id were found"))
	}
	return wishlist, nil
}
