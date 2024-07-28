package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type TradeProposalService struct {
	TradeProposalRepository repository.TradeProposalRepository
}

func NewTradeProposalService(TradeProposalRepository repository.TradeProposalRepository) *TradeProposalService {
	return &TradeProposalService{TradeProposalRepository: TradeProposalRepository}
}

func (service *TradeProposalService) GetAll() (*[]model.TradeProposal, error) {
	tradeProposals, err := service.TradeProposalRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade proposals were found"))
	}
	return &tradeProposals, nil
}

func (service *TradeProposalService) GetByID(id int) (*model.TradeProposal, error) {
	tradeProposal, err := service.TradeProposalRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade proposals with that id were found"))
	}
	return tradeProposal, nil
}

func (service *TradeProposalService) GetAllByTeamID(teamID int) (*[]model.TradeProposal, error) {
	tradeProposals, err := service.TradeProposalRepository.GetAllByTeamID(teamID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade proposals with that sender idwere found"))
	}
	return &tradeProposals, nil
}

func (service *TradeProposalService) Create(tradeProposal *model.TradeProposal) error {
	err := service.TradeProposalRepository.Create(tradeProposal)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no trade proposals were created"))
		return err
	}
	return nil
}

func (service *TradeProposalService) Update(tradeProposal *model.TradeProposal) error {
	err := service.TradeProposalRepository.Update(tradeProposal)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no trade proposals were updated"))
		return err
	}
	return nil
}

func (service *TradeProposalService) GetLatest() (*model.TradeProposal, error) {
	tradeProposal, err := service.TradeProposalRepository.GetLatest()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no latest trade proposals were found"))
	}
	return tradeProposal, nil
}
