package repository

import "basketball-league-server/model"

type TradeProposalRepository interface {
	GetAll() ([]model.TradeProposal, error)
	GetByID(id int) (*model.TradeProposal, error)
	GetAllReceivedByManagerID(managerID int) ([]model.TradeProposal, error)
	GetAllSentByManagerID(managerID int) ([]model.TradeProposal, error)
	Create(*model.TradeProposal) error
	Update(*model.TradeProposal) error
	GetLatest() (*model.TradeProposal, error)
}
