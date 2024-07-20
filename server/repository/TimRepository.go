package repository

import "basketball-league-server/model"

type TimRepository interface {
	GetAll() ([]model.Tim, error)
	GetByID(id int) (*model.Tim, error)
}
