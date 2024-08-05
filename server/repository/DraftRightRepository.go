package repository

import "basketball-league-server/model"

type DraftRightRepository interface {
	GetAll() ([]model.DraftRight, error)
	GetByID(id int) (*model.DraftRight, error)
	GetAllByTeamID(teamID int) ([]model.DraftRight, error)
	GetAllAvailableByTeamID(teamID int) ([]model.DraftRight, error)
	Update(draftRights *model.DraftRight) error
	AddToWishlist(draftRight *model.DraftRight, teamId int) error
	RemoveFromWishlist(draftRight *model.DraftRight, teamId int) error
}
