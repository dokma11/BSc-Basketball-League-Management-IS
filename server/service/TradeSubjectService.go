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
	TradeSubjects, err := service.TradeSubjectRepository.GetAll()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade subjects were found"))
	}
	return &TradeSubjects, nil
}

func (service *TradeSubjectService) GetByID(id int) (*model.TradeSubject, error) {
	TradeSubject, err := service.TradeSubjectRepository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade subjects with that id were found"))
	}
	return TradeSubject, nil
}

func (service *TradeSubjectService) GetAllByTradeProposalID(tradeProposalID int) (*[]model.TradeSubject, error) {
	TradeSubjects, err := service.TradeSubjectRepository.GetAllByTradeProposalID(tradeProposalID)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("no trade subjects with that trade proposal id were found"))
	}
	return &TradeSubjects, nil
}

func (service *TradeSubjectService) Create(TradeSubject *model.TradeSubject) error {
	err := service.TradeSubjectRepository.Create(TradeSubject)
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("no trade subjects were created"))
		return err
	}
	return nil
}
