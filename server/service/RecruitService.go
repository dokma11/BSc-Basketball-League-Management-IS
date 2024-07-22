package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type RecruitService struct {
	RecruitRepository repository.RecruitRepository
}

func NewRecruitService(RecruitRepository repository.RecruitRepository) *RecruitService {
	return &RecruitService{RecruitRepository: RecruitRepository}
}

func (service *RecruitService) GetAll() (*[]model.Recruit, error) {
	recruits, err := service.RecruitRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no recruits were found"))
	}

	return &recruits, nil
}

func (service *RecruitService) GetByID(id int) (*model.Recruit, error) {
	recruit, err := service.RecruitRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no recruits with that id were found"))
	}

	return recruit, nil
}

func (service *RecruitService) Create(recruit *model.Recruit) error {
	err := service.RecruitRepository.Create(recruit)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no recruits were created"))
		return err
	}
	return nil
}

func (service *RecruitService) Update(recruit *model.Recruit) error {
	err := service.RecruitRepository.Update(recruit)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no recruits were updated"))
		return err
	}
	return nil
}
