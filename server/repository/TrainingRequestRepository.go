package repository

import "basketball-league-server/model"

type TrainingRequestRepository interface {
	GetAll() ([]model.TrainingRequest, error)
	GetByID(id int) (*model.TrainingRequest, error)
	GetAllBySenderID(userID int) ([]model.TrainingRequest, error)
	GetAllByReceiverID(userID int) ([]model.TrainingRequest, error)
	Create(trainingRequest *model.TrainingRequest, recruitId int64) error
	Update(trainingRequest *model.TrainingRequest) error
}
