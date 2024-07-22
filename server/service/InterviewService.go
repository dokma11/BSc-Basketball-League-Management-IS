package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type InterviewService struct {
	InterviewRepository repository.InterviewRepository
}

func NewInterviewService(InterviewRepository repository.InterviewRepository) *InterviewService {
	return &InterviewService{InterviewRepository: InterviewRepository}
}

func (service *InterviewService) GetAll() (*[]model.Interview, error) {
	interviews, err := service.InterviewRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no interviews were found"))
	}
	return &interviews, nil
}

func (service *InterviewService) GetByID(id int) (*model.Interview, error) {
	interview, err := service.InterviewRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no interview with that id were found"))
	}
	return interview, nil
}

func (service *InterviewService) GetAllByUserID(userID int) (*[]model.Interview, error) {
	interviews, err := service.InterviewRepository.GetAllByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no interviews with that sender id were found"))
	}
	return &interviews, nil
}

func (service *InterviewService) Create(Interview *model.Interview) error {
	err := service.InterviewRepository.Create(Interview)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no interviews were created"))
		return err
	}
	return nil
}
