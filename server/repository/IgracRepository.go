package repository

import (
	"basketball-league-server/model"
)

type IgracRepository interface {
	GetAll() ([]model.Igrac, error)
	GetByID(id int) (*model.Igrac, error)
	GetAllByTeamID(teamId int) ([]model.Igrac, error)
}
