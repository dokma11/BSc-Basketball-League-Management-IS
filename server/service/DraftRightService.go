package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
	"log"
)

type DraftRightService struct {
	DraftRightRepository repository.DraftRightRepository
}

func NewDraftRightService(DraftRightRepository repository.DraftRightRepository) *DraftRightService {
	return &DraftRightService{DraftRightRepository: DraftRightRepository}
}

func (service *DraftRightService) GetAll() (*[]model.DraftRight, error) {
	draftRights, err := service.DraftRightRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no draft rights were found"))
	}

	return &draftRights, nil
}

func (service *DraftRightService) GetByID(id int) (*model.DraftRight, error) {
	draftRight, err := service.DraftRightRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no draft rights with that id were found"))
	}

	return draftRight, nil
}

func (service *DraftRightService) GetAllByTeamID(teamID int) ([]model.DraftRight, error) {
	draftRights, err := service.DraftRightRepository.GetAllByTeamID(teamID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no draft rights with that team id were found"))
	}

	return draftRights, nil
}

func (service *DraftRightService) GetAllAvailableByTeamID(teamID int) ([]model.DraftRight, error) {
	draftRights, err := service.DraftRightRepository.GetAllAvailableByTeamID(teamID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no available draft rights with that team id were found"))
	}

	return draftRights, nil
}

func (service *DraftRightService) Update(draftRights *model.DraftRight) error {
	err := service.DraftRightRepository.Update(draftRights)
	if err != nil {
		log.Println("Error updating draft rights")
		return err
	}
	return nil
}

func (service *DraftRightService) AddToWishlist(draftRights *model.DraftRight, teamId int) error {
	err := service.DraftRightRepository.AddToWishlist(draftRights, teamId)
	if err != nil {
		log.Println("Error adding draft rights to the wishlist")
		return err
	}
	return nil
}

func (service *DraftRightService) RemoveFromWishlist(draftRights *model.DraftRight, teamId int) error {
	err := service.DraftRightRepository.RemoveFromWishlist(draftRights, teamId)
	if err != nil {
		log.Println("Error removing draft rights from the wishlist")
		return err
	}
	return nil
}
