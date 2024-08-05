package repository

import "basketball-league-server/model"

type TeamRepository interface {
	GetAll() ([]model.Team, error)
	GetByID(id int) (*model.Team, error)
	GetByUserID(userID int) (*model.Team, error)
	GetPlayerTradeDestination(userID int) (*model.Team, error)
	GetPickTradeDestination(userID int) (*model.Team, error)
	GetDraftRightsTradeDestination(userID int) (*model.Team, error)
	GetWishlistByTeamID(teamID int) ([]model.WishlistAsset, error)
}
