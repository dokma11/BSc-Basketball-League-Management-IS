package model

import (
	"errors"
	"time"
)

type TradeProposalStatus int

const (
	IN_PROGRESS TradeProposalStatus = iota
	ACCEPTED
	DECLINED
	CANCELLED
)

type TradeType int

const (
	PLAYER_PLAYER TradeType = iota
	PLAYER_PICK
	PICK_PICK
)

type TradeProposal struct {
	ID           int64               `json:"idZahTrg"`
	Date         time.Time           `json:"datZahTrg"` // Date of creation
	Type         TradeType           `json:"tipZahTrg"`
	Status       TradeProposalStatus `json:"statusZahTrg"`
	DenialReason string              `json:"razlogOdbij"`    // Reason for declining if declined
	SenderId     int64               `json:"idMenadzerPos"`  // Sender foreign key
	ReceiverId   int64               `json:"idMenadzerPrim"` // Receiver foreign key
}

func NewTradeProposal(idZahTrg int64, datZahTrg time.Time, tipZahTrg TradeType, statusZahTrg TradeProposalStatus,
	razlogOdbij string) (*TradeProposal, error) {
	tradeProposal := &TradeProposal{
		ID:           idZahTrg,
		Date:         datZahTrg,
		Type:         tipZahTrg,
		Status:       statusZahTrg,
		DenialReason: razlogOdbij,
	}

	if err := tradeProposal.Validate(); err != nil {
		return nil, err
	}

	return tradeProposal, nil
}

func (t *TradeProposal) Validate() error {
	if t.Type < 0 || t.Type > 2 {
		return errors.New("type field is invalid")
	}
	if t.Status < 0 || t.Status > 3 {
		return errors.New("status field is invalid")
	}
	if t.Status == 2 && t.DenialReason == "" {
		return errors.New("reason for declining field can not be empty when the proposal is declined")
	}
	return nil
}

type TradeProposalDAO struct {
	IdZahTrg       int64               `json:"idZahTrg"`
	DatZahTrg      time.Time           `json:"datZahTrg"` // Date of creation
	TipZahTrg      TradeType           `json:"tipZahTrg"`
	StatusZahTrg   TradeProposalStatus `json:"statusZahTrg"`
	RazlogOdbij    string              `json:"razlogOdbij"`    // Reason for declining if declined
	IdMenadzerPos  int64               `json:"idMenadzerPos"`  // Sender foreign key
	IdMenadzerPrim int64               `json:"idMenadzerPrim"` // Receiver foreign key
}

func (t *TradeProposal) FromDAO(tradeProposalDAO *TradeProposalDAO) {
	t.ID = tradeProposalDAO.IdZahTrg
	t.Date = tradeProposalDAO.DatZahTrg
	t.Type = tradeProposalDAO.TipZahTrg
	t.Status = tradeProposalDAO.StatusZahTrg
	t.DenialReason = tradeProposalDAO.RazlogOdbij
	t.SenderId = tradeProposalDAO.IdMenadzerPos
	t.ReceiverId = tradeProposalDAO.IdMenadzerPrim
}

type TradeProposalCreateDTO struct {
	DatZahTrg         time.Time `json:"datZahTrg"` // Date of creation
	TipZahTrg         TradeType `json:"tipZahTrg"`
	IdMenadzerPos     int64     `json:"idMenadzerPos"`     // Sender foreign key
	IdMenadzerPrimTim int64     `json:"idMenadzerPrimTim"` // Receiver team foreign key
}

type TradeProposalUpdateDTO struct {
	IdZahTrg  int64     `json:"idZahTrg"`
	TipZahTrg TradeType `json:"tipZahTrg"`
}
