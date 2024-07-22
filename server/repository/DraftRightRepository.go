package repository

import "basketball-league-server/model"

type DraftRightRepository interface {
	GetAll() ([]model.DraftRight, error)
	GetByID(id int) (*model.DraftRight, error)
}
