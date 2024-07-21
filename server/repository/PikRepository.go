package repository

import "basketball-league-server/model"

type PikRepository interface {
	GetAll() ([]model.Pik, error)
	GetByID(id int) (*model.Pik, error)
	GetAllByTeamID(teamId int) ([]model.Pik, error)
	GetAllByYear(year string) ([]model.Pik, error)
}
