package model

import (
	"errors"
	"time"
)

type Trade struct {
	IdTrg    int64     `json:"idTrg"`
	DatTrg   time.Time `json:"datTrg"` // Date of trade occurrence
	TipTrg   TradeType `json:"tipTrg"`
	IdZahTrg int64     `json:"idZahTrg"` // Trade Proposal foreign key
}

func NewTrade(datTrg time.Time, tipTrg TradeType, idZahTrg int64) (*Trade, error) {
	trade := &Trade{
		DatTrg:   datTrg,
		TipTrg:   tipTrg,
		IdZahTrg: idZahTrg,
	}

	if err := trade.Validate(); err != nil {
		return nil, err
	}

	return trade, nil
}

func (t *Trade) Validate() error {
	if t.TipTrg < 0 || t.TipTrg > 2 {
		return errors.New("trade type field is invalid")
	}
	return nil
}
