package service

import (
	"basketball-league-server/model"
	"basketball-league-server/repository"
	"fmt"
)

type TradeSubjectService struct {
	TradeSubjectRepository repository.TradeSubjectRepository
}

func NewTradeSubjectService(TradeSubjectRepository repository.TradeSubjectRepository) *TradeSubjectService {
	return &TradeSubjectService{TradeSubjectRepository: TradeSubjectRepository}
}

func (service *TradeSubjectService) GetAll() (*[]model.TradeSubject, error) {
	tradeSubjects, err := service.TradeSubjectRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade subjects were found"))
	}
	return &tradeSubjects, nil
}

func (service *TradeSubjectService) GetByID(id int) (*model.TradeSubject, error) {
	tradeSubject, err := service.TradeSubjectRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade subjects with that id were found"))
	}
	return tradeSubject, nil
}

func (service *TradeSubjectService) GetAllByTradeProposalID(tradeProposalID int) (*[]model.TradeSubject, error) {
	tradeSubjects, err := service.TradeSubjectRepository.GetAllByTradeProposalID(tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade subjects with that trade proposal id were found"))
	}
	return &tradeSubjects, nil
}

func (service *TradeSubjectService) GetPlayerTypeSubjectsByTradeProposalID(tradeProposalID int) (*[]model.TradeSubjectDetailsResponseDTO, error) {
	tradeSubjects, err := service.TradeSubjectRepository.GetPlayerTypeSubjectsByTradeProposalID(tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no player type trade subjects with that trade proposal id were found"))
	}
	return &tradeSubjects, nil
}

func (service *TradeSubjectService) GetPickTypeSubjectsByTradeProposalID(tradeProposalID int) (*[]model.TradeSubjectDetailsResponseDTO, error) {
	tradeSubjects, err := service.TradeSubjectRepository.GetPickTypeSubjectsByTradeProposalID(tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no pick type trade subjects with that trade proposal id were found"))
	}
	return &tradeSubjects, nil
}

func (service *TradeSubjectService) GetDraftRightsTypeSubjectsByTradeProposalID(tradeProposalID int) (*[]model.TradeSubjectDetailsResponseDTO, error) {
	tradeSubjects, err := service.TradeSubjectRepository.GetDraftRightsTypeSubjectsByTradeProposalID(tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no draft rights type trade subjects with that trade proposal id were found"))
	}
	return &tradeSubjects, nil
}

func (service *TradeSubjectService) Create(TradeSubject *model.TradeSubject) error {
	err := service.TradeSubjectRepository.Create(TradeSubject)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no trade subjects were created"))
		return err
	}
	return nil
}
