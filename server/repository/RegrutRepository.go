package repository

import "basketball-league-server/model"

type RegrutRepository interface {
	GetAll() ([]model.Regrut, error)
	GetByID(id int) (*model.Regrut, error)
	Create(*model.Regrut) error
	Update(*model.Regrut) error
}
