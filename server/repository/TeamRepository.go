package repository

import "basketball-league-server/model"

type TeamRepository interface {
	GetAll() ([]model.Team, error)
	GetByID(id int) (*model.Team, error)
	GetByUserID(userID int) (*model.Team, error)
}
