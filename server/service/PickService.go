package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
	"log"
)

type PickService struct {
	PickRepository repository.PickRepository
}

func NewPickService(pikRepository repository.PickRepository) *PickService {
	return &PickService{PickRepository: pikRepository}
}

func (service *PickService) GetAll() (*[]model.Pick, error) {
	picks, err := service.PickRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks were found"))
	}

	return &picks, nil
}

func (service *PickService) GetByID(id int) (*model.Pick, error) {
	pick, err := service.PickRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks with that id were found"))
	}

	return pick, nil
}

func (service *PickService) GetAllByTeamID(teamId int) (*[]model.Pick, error) {
	picks, err := service.PickRepository.GetAllByTeamID(teamId)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks with that team id were found"))
	}

	return &picks, nil
}

func (service *PickService) GetAllByYear(year string) (*[]model.Pick, error) {
	picks, err := service.PickRepository.GetAllByYear(year)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no picks in that year were found"))
	}

	return &picks, nil
}

func (service *PickService) Update(pick *model.Pick) error {
	err := service.PickRepository.Update(pick)
	if err != nil {
		log.Println("Error updating pick")
		return err
	}
	return nil
}
