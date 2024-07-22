package repository

import "basketball-league-server/model"

type TrainingRepository interface {
	GetAll() ([]model.Training, error)
	GetByID(id int) (*model.Training, error)
	GetAllByUserID(userID int) ([]model.Training, error)
	Create(*model.Training) error
}
