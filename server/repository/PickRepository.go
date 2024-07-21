package repository

import "basketball-league-server/model"

type PickRepository interface {
	GetAll() ([]model.Pick, error)
	GetByID(id int) (*model.Pick, error)
	GetAllByTeamID(teamId int) ([]model.Pick, error)
	GetAllByYear(year string) ([]model.Pick, error)
}