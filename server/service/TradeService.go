package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type TradeService struct {
	TradeRepository repository.TradeRepository
}

func NewTradeService(TradeRepository repository.TradeRepository) *TradeService {
	return &TradeService{TradeRepository: TradeRepository}
}

func (service *TradeService) GetAll() (*[]model.Trade, error) {
	trades, err := service.TradeRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade were found"))
	}
	return &trades, nil
}

func (service *TradeService) GetByID(id int) (*model.Trade, error) {
	trade, err := service.TradeRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade with that id were found"))
	}
	return trade, nil
}

func (service *TradeService) GetAllByTeamID(teamID int) (*[]model.Trade, error) {
	trades, err := service.TradeRepository.GetAllByTeamID(teamID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade with that sender idwere found"))
	}
	return &trades, nil
}

func (service *TradeService) Create(Trade *model.Trade) error {
	err := service.TradeRepository.Create(Trade)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no trade were created"))
		return err
	}
	return nil
}
