package repository

import "basketball-league-server/model"

type TradeRepository interface {
	GetAll() ([]model.Trade, error)
	GetByID(id int) (*model.Trade, error)
	GetAllByTeamID(teamId int) ([]model.Trade, error)
	Create(*model.Trade) error
}
