package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type TrainingService struct {
	TrainingRepository repository.TrainingRepository
}

func NewTrainingService(TrainingRepository repository.TrainingRepository) *TrainingService {
	return &TrainingService{TrainingRepository: TrainingRepository}
}

func (service *TrainingService) GetAll() (*[]model.Training, error) {
	trainings, err := service.TrainingRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trainings were found"))
	}
	return &trainings, nil
}

func (service *TrainingService) GetByID(id int) (*model.Training, error) {
	training, err := service.TrainingRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no training with that id were found"))
	}
	return training, nil
}

func (service *TrainingService) GetAllByUserID(userID int) (*[]model.Training, error) {
	Trainings, err := service.TrainingRepository.GetAllByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trainings with that sender id were found"))
	}
	return &Trainings, nil
}

func (service *TrainingService) Create(Training *model.Training) error {
	err := service.TrainingRepository.Create(Training)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no trainings were created"))
		return err
	}
	return nil
}
