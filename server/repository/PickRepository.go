package repository

import "basketball-league-server/model"

type PickRepository interface {
	GetAll() ([]model.Pick, error)
	GetByID(id int) (*model.Pick, error)
	GetAllByTeamID(teamId int) ([]model.Pick, error)
	GetAllAvailableByTeamID(teamId int) ([]model.Pick, error)
	GetAllByYear(year string) ([]model.Pick, error)
	Update(pick *model.Pick) error
	AddToWishlist(pick *model.Pick, teamId int) error
	RemoveFromWishlist(pick *model.Pick, teamId int) error
}
