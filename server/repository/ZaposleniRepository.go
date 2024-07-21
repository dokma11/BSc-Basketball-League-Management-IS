package repository

import "basketball-league-server/model"

type ZaposleniRepository interface {
	GetAll() ([]model.Zaposleni, error)
	GetByID(id int) (*model.Zaposleni, error)
}
