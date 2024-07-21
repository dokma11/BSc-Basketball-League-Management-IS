package repository

import (
	"basketball-league-server/model"
)

type PlayerRepository interface {
	GetAll() ([]model.Player, error)
	GetByID(id int) (*model.Player, error)
	GetAllByTeamID(teamId int) ([]model.Player, error)
}
