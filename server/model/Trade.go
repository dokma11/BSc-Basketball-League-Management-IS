package model

import (
	"errors"
	"time"
)

type Trade struct {
	ID              int64     `json:"idTrg"`
	OccurrenceDate  time.Time `json:"datTrg"` // Date of trade occurrence
	Type            TradeType `json:"tipTrg"`
	TradeProposalId int64     `json:"idZahTrg"` // Trade Proposal foreign key
}

func NewTrade(datTrg time.Time, tipTrg TradeType, idZahTrg int64) (*Trade, error) {
	trade := &Trade{
		OccurrenceDate:  datTrg,
		Type:            tipTrg,
		TradeProposalId: idZahTrg,
	}

	if err := trade.Validate(); err != nil {
		return nil, err
	}

	return trade, nil
}

func (t *Trade) Validate() error {
	if t.Type < 0 || t.Type > 2 {
		return errors.New("trade type field is invalid")
	}
	return nil
}

type TradeDAO struct {
	IdTrg    int64     `json:"idTrg"`
	DatTrg   time.Time `json:"datTrg"` // Date of trade occurrence
	TipTrg   TradeType `json:"tipTrg"`
	IdZahTrg int64     `json:"idZahTrg"` // Trade Proposal foreign key
}

func (t *Trade) FromDAO(tradeDAO *TradeDAO) {
	t.ID = tradeDAO.IdTrg
	t.OccurrenceDate = tradeDAO.DatTrg
	t.Type = tradeDAO.TipTrg
	t.TradeProposalId = tradeDAO.IdZahTrg
}
