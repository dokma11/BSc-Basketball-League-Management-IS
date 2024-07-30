package repository

import "basketball-league-server/model"

type TradeSubjectRepository interface {
	GetAll() ([]model.TradeSubject, error)
	GetByID(id int) (*model.TradeSubject, error)
	GetAllByTradeProposalID(tradeProposalID int) ([]model.TradeSubject, error)
	Create(*model.TradeSubject) error
	GetPlayerTypeSubjectsByTradeProposalID(tradeProposalID int) ([]model.TradeSubjectDetailsResponseDTO, error)
	GetPickTypeSubjectsByTradeProposalID(tradeProposalID int) ([]model.TradeSubjectDetailsResponseDTO, error)
	GetDraftRightsTypeSubjectsByTradeProposalID(tradeProposalID int) ([]model.TradeSubjectDetailsResponseDTO, error)
}
