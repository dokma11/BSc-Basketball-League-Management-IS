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
	IdZahTrg     int64               `json:"idZahTrg"`
	DatZahTrg    time.Time           `json:"datZahTrg"` // Date of creation
	TipZahTrg    TradeType           `json:"tipZahTrg"`
	StatusZahTrg TradeProposalStatus `json:"statusZahTrg"`
	RazlogOdbij  string              `json:"razlogOdbij"` // Reason for declining if declined
}

func NewTradeProposal(idZahTrg int64, datZahTrg time.Time, tipZahTrg TradeType, statusZahTrg TradeProposalStatus,
	razlogOdbij string) (*TradeProposal, error) {
	tradeProposal := &TradeProposal{
		IdZahTrg:     idZahTrg,
		DatZahTrg:    datZahTrg,
		TipZahTrg:    tipZahTrg,
		StatusZahTrg: statusZahTrg,
		RazlogOdbij:  razlogOdbij,
	}

	if err := tradeProposal.Validate(); err != nil {
		return nil, err
	}

	return tradeProposal, nil
}

func (t *TradeProposal) Validate() error {
	if t.TipZahTrg < 0 || t.TipZahTrg > 2 {
		return errors.New("type field is invalid")
	}
	if t.StatusZahTrg < 0 || t.StatusZahTrg > 3 {
		return errors.New("status field is invalid")
	}
	if t.StatusZahTrg == 2 && t.RazlogOdbij == "" {
		return errors.New("reason for declining field can not be empty when the proposal is declined")
	}

	return nil
}
