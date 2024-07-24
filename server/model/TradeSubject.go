package model

import "errors"

type TradeSubjectType int

const (
	Igrac         TradeSubjectType = iota // Player
	Pik                                   // Pick
	PravaNaIgraca                         // Draft Rights
)

type TradeSubject struct {
	IdPredTrg  int64            `json:"idPredTrg"`
	TipPredTrg TradeSubjectType `json:"tipPredTrg"`
	IdPrava    int64            `json:"idPrava"`  // Draft Rights foreign key
	IdIgrac    int64            `json:"idIgrac"`  // Player foreign key
	IdZahTrg   int64            `json:"idZahTrg"` // Trade Request foreign key
	IdPik      int64            `json:"idPik"`    // Pick foreign key
}

func NewTradeSubject(idPredTrg int64, tipPredTrg TradeSubjectType) (*TradeSubject, error) {
	tradeSubject := &TradeSubject{
		IdPredTrg:  idPredTrg,
		TipPredTrg: tipPredTrg,
	}

	if err := tradeSubject.Validate(); err != nil {
		return nil, err
	}

	return tradeSubject, nil
}

func (t *TradeSubject) Validate() error {
	if t.TipPredTrg < 0 || t.TipPredTrg > 2 {
		return errors.New("trade subject type field is invalid")
	}
	return nil
}
