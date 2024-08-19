package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type DraftService struct {
	DraftRepository repository.DraftRepository
}

func NewDraftService(DraftRepository repository.DraftRepository) *DraftService {
	return &DraftService{DraftRepository: DraftRepository}
}

func (service *DraftService) GetAll() (*[]model.Draft, error) {
	drafts, err := service.DraftRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no drafts were found"))
	}

	return &drafts, nil
}

func (service *DraftService) GetByID(id int) (*model.Draft, error) {
	draft, err := service.DraftRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no drafts with that id were found"))
	}

	return draft, nil
}

func (service *DraftService) GetLatest() (*model.Draft, error) {
	draft, err := service.DraftRepository.GetLatest()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no latest drafts were found"))
	}

	return draft, nil
}
