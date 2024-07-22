package model

import "errors"

type TradeSubject struct {
	IdPredTrg     int64  `json:"idPredTrg"`
	KomentPredTrg string `json:"komentPredTrg"` // Comment about the trade subject
}

func NewTradeSubject(idPredTrg int64, komentPredTrg string) (*TradeSubject, error) {
	tradeSubject := &TradeSubject{
		IdPredTrg:     idPredTrg,
		KomentPredTrg: komentPredTrg,
	}

	if err := tradeSubject.Validate(); err != nil {
		return nil, err
	}

	return tradeSubject, nil
}

func (t *TradeSubject) Validate() error {
	if t.KomentPredTrg == "" {
		return errors.New("comment field is invalid")
	}
	return nil
}
