package repository

import "basketball-league-server/model"

type DraftRepository interface {
	GetAll() ([]model.Draft, error)
	GetByID(id int) (*model.Draft, error)
}
