package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type InterviewRequestService struct {
	InterviewRequestRepository repository.InterviewRequestRepository
}

func NewInterviewRequestService(InterviewRequestRepository repository.InterviewRequestRepository) *InterviewRequestService {
	return &InterviewRequestService{InterviewRequestRepository: InterviewRequestRepository}
}

func (service *InterviewRequestService) GetAll() (*[]model.InterviewRequest, error) {
	interviewRequests, err := service.InterviewRequestRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no interview requests were found"))
	}
	return &interviewRequests, nil
}

func (service *InterviewRequestService) GetByID(id int) (*model.InterviewRequest, error) {
	interviewRequest, err := service.InterviewRequestRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no interview requests with that id were found"))
	}
	return interviewRequest, nil
}

func (service *InterviewRequestService) GetAllBySenderID(userID int) (*[]model.InterviewRequest, error) {
	interviewRequests, err := service.InterviewRequestRepository.GetAllBySenderID(userID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no interview requests with that sender id were found"))
	}
	return &interviewRequests, nil
}

func (service *InterviewRequestService) GetAllByReceiverID(userID int) (*[]model.InterviewRequest, error) {
	interviewRequests, err := service.InterviewRequestRepository.GetAllByReceiverID(userID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no interview requests with that receiver id were found"))
	}
	return &interviewRequests, nil
}

func (service *InterviewRequestService) Create(InterviewRequest *model.InterviewRequest) error {
	err := service.InterviewRequestRepository.Create(InterviewRequest)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no interview requests were created"))
		return err
	}
	return nil
}

func (service *InterviewRequestService) Update(InterviewRequest *model.InterviewRequest) error {
	err := service.InterviewRequestRepository.Update(InterviewRequest)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no interview requests were updated"))
		return err
	}
	return nil
}
