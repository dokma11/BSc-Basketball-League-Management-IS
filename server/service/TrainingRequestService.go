package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type TrainingRequestService struct {
	TrainingRequestRepository repository.TrainingRequestRepository
}

func NewTrainingRequestService(TrainingRequestRepository repository.TrainingRequestRepository) *TrainingRequestService {
	return &TrainingRequestService{TrainingRequestRepository: TrainingRequestRepository}
}

func (service *TrainingRequestService) GetAll() (*[]model.TrainingRequest, error) {
	trainingRequests, err := service.TrainingRequestRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no training requests were found"))
	}
	return &trainingRequests, nil
}

func (service *TrainingRequestService) GetByID(id int) (*model.TrainingRequest, error) {
	trainingRequest, err := service.TrainingRequestRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no training requests with that id were found"))
	}
	return trainingRequest, nil
}

func (service *TrainingRequestService) GetAllBySenderID(userID int) (*[]model.TrainingRequest, error) {
	trainingRequests, err := service.TrainingRequestRepository.GetAllBySenderID(userID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no training requests with that sender id were found"))
	}
	return &trainingRequests, nil
}

func (service *TrainingRequestService) GetAllByReceiverID(userID int) (*[]model.TrainingRequest, error) {
	trainingRequests, err := service.TrainingRequestRepository.GetAllByReceiverID(userID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no training requests with that receiver id were found"))
	}
	return &trainingRequests, nil
}

func (service *TrainingRequestService) Create(TrainingRequest *model.TrainingRequest) error {
	err := service.TrainingRequestRepository.Create(TrainingRequest)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no training requests were created"))
		return err
	}
	return nil
}

func (service *TrainingRequestService) Update(TrainingRequest *model.TrainingRequest) error {
	err := service.TrainingRequestRepository.Update(TrainingRequest)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no training requests were updated"))
		return err
	}
	return nil
}
